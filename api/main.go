package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rparaschak/mono-tmpl/api/pkg/app"
	"github.com/rparaschak/mono-tmpl/api/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	application := app.New(cfg)

	go func() {
		slog.Info("starting server", "addr", application.Server.Addr, "env", cfg.HTTPServer.Env)
		if err := application.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := application.Server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}
}
