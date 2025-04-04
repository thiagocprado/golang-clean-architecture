package main

import (
	"fmt"
	"log/slog"

	"github.com/thiagocprado/golang-api-structure/internal/database"
	"github.com/thiagocprado/golang-api-structure/internal/dependencies"
	"github.com/thiagocprado/golang-api-structure/internal/env"
	"github.com/thiagocprado/golang-api-structure/pkg/logger"
	"github.com/thiagocprado/golang-api-structure/pkg/middlewares"
	"github.com/thiagocprado/golang-api-structure/pkg/router"

	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	logger.InitSlog()
	env.LoadEnvs()

	err := database.NewMySqlConn()
	if err != nil {
		os.Exit(-1)
	}

	startHttpServer()
}

func startHttpServer() {
	r := router.New()
	dependencies.LoadDependencies(r)

	addr := fmt.Sprintf(":%d", env.HttpPort)
	c := middlewares.CORS()
	server := &http.Server{
		Addr:         addr,
		Handler:      c.Handler(r.Router),
		ReadTimeout:  time.Duration(env.HttpReadTimeout) * time.Second,
		WriteTimeout: time.Duration(env.HttpWriteTimeout) * time.Second,
	}

	slog.Info(
		"Starting HTTP server",
		slog.String("addr", strconv.Itoa(env.HttpPort)),
		slog.Int("read_timeout", env.HttpReadTimeout),
		slog.Int("write_timeout", env.HttpWriteTimeout),
	)

	if err := server.ListenAndServe(); err != nil {
		slog.Error(
			"Failed to start HTTP server",
			slog.String("error", err.Error()),
			slog.String("addr", addr),
			slog.Int("read_timeout", env.HttpReadTimeout),
			slog.Int("write_timeout", env.HttpWriteTimeout),
		)
		os.Exit(-1)
	}
}
