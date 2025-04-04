package requests

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"github.com/thiagocprado/golang-api-structure/pkg/errs"
	"github.com/thiagocprado/golang-api-structure/pkg/handles"

	"net/http"
	"net/url"
	"strings"
)

type CustomTransport struct {
	Token string
	Base  http.RoundTripper
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token := t.Token
	if !strings.Contains(token, "Bearer") {
		token = "Bearer " + token
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-type", "application/json")

	resp, err := t.Base.RoundTrip(req)
	if err != nil {
		slog.Error("Falha ao fazer requisição!", slog.String("message", err.Error()))
		return nil, err
	}

	return resp, nil
}

func (api *ApiConfig) Get(path string, params map[string]string) (*http.Response, *errs.Err) {
	urlPath := api.BaseURL + path
	if len(params) > 0 {
		query := url.Values{}
		for key, value := range params {
			query.Add(key, value)
		}
		urlPath += "?" + query.Encode()
	}

	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, errs.InternalServerError("erro ao criar requisição GET", err)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, errs.InternalServerError("erro ao fazer requisição GET", err)
	}

	return handles.Response(resp)
}

func (api *ApiConfig) Post(path string, body any) (*http.Response, *errs.Err) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, errs.InternalServerError("erro ao serializar body", err)
	}

	req, err := http.NewRequest("POST", api.BaseURL+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, errs.InternalServerError("erro ao criar requisição POST", err)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, errs.InternalServerError("erro ao fazer requisição POST", err)
	}

	return handles.Response(resp)
}

func (api *ApiConfig) Put(path string, body any) (*http.Response, *errs.Err) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, errs.InternalServerError("erro ao serializar body", err)
	}

	req, err := http.NewRequest("PUT", api.BaseURL+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, errs.InternalServerError("erro ao criar requisição PUT", err)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, errs.InternalServerError("erro ao fazer requisição PUT", err)
	}

	return handles.Response(resp)
}

func (api *ApiConfig) Delete(path string, params map[string]string) (*http.Response, *errs.Err) {
	urlPath := api.BaseURL + path
	if len(params) > 0 {
		query := url.Values{}
		for key, value := range params {
			query.Add(key, value)
		}
		urlPath += "?" + query.Encode()
	}

	req, err := http.NewRequest("DELETE", urlPath, nil)
	if err != nil {
		return nil, errs.InternalServerError("erro ao criar requisição DELETE", err)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, errs.InternalServerError("erro ao fazer requisição DELETE", err)
	}

	return handles.Response(resp)
}
