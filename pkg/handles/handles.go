package handles

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/thiagocprado/golang-api-structure/internal/env"
	"github.com/thiagocprado/golang-api-structure/pkg/errs"

	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type AppInfo struct {
	Env         string `json:"env"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type ResponseBody struct {
	AppInfo AppInfo `json:"app_info"`
	Data    any     `json:"data"`
}

func Error(w http.ResponseWriter, e *errs.Err) {
	slog.Error(
		e.Message,
		slog.String("error", e.Error),
		slog.Int("status", e.StatusCode),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)

	if err := json.NewEncoder(w).Encode(e); err != nil {
		slog.Error("failed to encode response", slog.String("error", err.Error()))
		json.NewEncoder(w).Encode(err.Error())
	}
}

func Payload[T any](body io.ReadCloser) (T, *errs.Err) {
	var result T

	err := json.NewDecoder(body).Decode(&result)
	if err != nil {
		slog.Error("failed to decode payload", slog.String("error", err.Error()))
		return result, nil
	}

	err = validator.New().Struct(result)
	if err != nil {
		errors := err.(validator.ValidationErrors)

		var errMsgs []string
		for _, e := range errors {
			field := strings.Split(e.Namespace(), ".")[1]

			errMsgs = append(errMsgs, field)
		}

		errMsg := fmt.Sprintf("O(s) campo(s) %s está(ão) inválido(s), verifique e tente novamente", strings.Join(errMsgs, ", "))

		return result, errs.BadRequest(errMsg, err)
	}

	return result, nil
}

func Request[T any](url url.Values) (T, *errs.Err) {
	var result T

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := decoder.Decode(&result, url)
	if err != nil {
		slog.Error("failed to decode request parameters", slog.String("error", err.Error()))
		return result, errs.BadRequest("Invalid request parameters", err)
	}

	if err := validator.New().Struct(result); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			return result, errs.BadRequest("Erro de validação desconhecido", err)
		}

		var errMsgs []string
		for _, e := range errors {
			field := strings.Split(e.Namespace(), ".")[1]
			errMsgs = append(errMsgs, field)
		}

		errMsg := fmt.Sprintf("O(s) campo(s) %s está(ão) inválido(s), verifique e tente novamente", strings.Join(errMsgs, ", "))
		return result, errs.BadRequest(errMsg, err)
	}

	return result, nil
}

func Response(resp *http.Response) (*http.Response, *errs.Err) {
	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()

		var apiError struct {
			Message string `json:"message"`
		}

		err := json.NewDecoder(resp.Body).Decode(&apiError)
		if err != nil {
			slog.Error(apiError.Message, slog.Int("status", resp.StatusCode))
		}

		switch resp.StatusCode {
		case http.StatusBadRequest:
			if apiError.Message == "" {
				apiError.Message = "Erro no payload!"
			}
			return nil, errs.BadRequest(apiError.Message, errors.New(apiError.Message))
		case http.StatusUnauthorized:
			if apiError.Message == "" {
				apiError.Message = "Não autorizado!"
			}
			return nil, errs.Unauthorized(apiError.Message, errors.New(apiError.Message))
		case http.StatusForbidden:
			if apiError.Message == "" {
				apiError.Message = "Permissão negada!"
			}
			return nil, errs.Forbidden(apiError.Message, errors.New(apiError.Message))
		case http.StatusNotFound:
			if apiError.Message == "" {
				apiError.Message = "Não encontrado!"
			}
			return nil, errs.NotFound(apiError.Message, errors.New(apiError.Message))
		case http.StatusConflict:
			if apiError.Message == "" {
				apiError.Message = "Conflito!"
			}
			return nil, errs.Conflict(apiError.Message, errors.New(apiError.Message))
		case http.StatusBadGateway:
			if apiError.Message == "" {
				apiError.Message = "Erro na comunicação com o servidor!"
			}
			return nil, errs.BadGateway(apiError.Message, errors.New(apiError.Message))
		case http.StatusServiceUnavailable:
			if apiError.Message == "" {
				apiError.Message = "Serviço indisponível!"
			}
			return nil, errs.ServiceUnavailable(apiError.Message, errors.New(apiError.Message))
		case http.StatusGatewayTimeout:
			if apiError.Message == "" {
				apiError.Message = "Tempo de resposta excedido!"
			}
			return nil, errs.GatewayTimeout(apiError.Message, errors.New(apiError.Message))

		default:
			if apiError.Message == "" {
				apiError.Message = "Erro interno no servidor!"
			}
			return nil, errs.InternalServerError(apiError.Message, errors.New(apiError.Message))
		}
	}

	return resp, nil
}

func Success(w http.ResponseWriter, statusCode int, data any) {
	responseBody := ResponseBody{
		Data: data,
		AppInfo: AppInfo{
			Env:         env.Environment,
			Name:        "solicitadores-budget-api",
			Version:     env.Version,
			Description: "Api criada com o intuíto de lida com solicitações de Verbas",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		Error(w, errs.InternalServerError("failed to encode response", err))
	}
}
