package middlewares

import (
	"github.com/thiagocprado/golang-api-structure/internal/env"

	"github.com/rs/cors"
)

func CORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{env.CorsAllowedOrigins}, // All origins
		AllowedMethods: []string{"POST", "GET", "PUT"},   // Allowing only get, just an example
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization", "X-Requested-With"},
	})
}
