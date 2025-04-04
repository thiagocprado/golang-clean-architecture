package middlewares

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/thiagocprado/golang-api-structure/internal/env"
	"github.com/thiagocprado/golang-api-structure/pkg/errs"
	"github.com/thiagocprado/golang-api-structure/pkg/handles"

	"github.com/golang-jwt/jwt/v5"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	auth := func(w http.ResponseWriter, r *http.Request) {
		if err := verifyToken(r); err != nil {
			handles.Error(w, err)
			return
		}

		next(w, r)
	}

	return auth
}

func verifyToken(r *http.Request) *errs.Err {
	auth := r.Header.Get("Authorization")

	if auth == "" {
		slog.Warn(
			"Invalid authentication token",
			slog.String("method", r.Method),
			slog.String("request_uri", r.RequestURI),
			slog.String("host", r.Host),
		)

		return errs.Unauthorized("Token inválido!", errors.New("invalid token"))
	}

	token := strings.Split(auth, " ")[1]
	jwtToken, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		return []byte(env.JwtSecretKey), nil
	})

	if err != nil {
		slog.Error("Falha ao fazer o parse do token!", slog.String("err", err.Error()), slog.String("token", token))
		return errs.Unauthorized("Token inválido!", errors.New("failed to parse token"))
	}

	if !jwtToken.Valid {
		slog.Error("Não foi possível validar o token!", slog.String("token", token))
		return errs.Unauthorized("Token inválido!", errors.New("failed to validate token"))
	}

	return nil
}
