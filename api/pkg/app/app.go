package app

import (
	"fmt"
	"net/http"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
)

type App struct {
	Config config.Config
	Server *http.Server
}

func New(cfg config.Config) *App {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	return &App{
		Config: cfg,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.HTTPServer.Port),
			Handler: mux,
		},
	}
}
