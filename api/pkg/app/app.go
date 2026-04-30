package app

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
	"github.com/rparaschak/mono-tmpl/api/pkg/httpapi"
)

type App struct {
	Config config.Config
	API    huma.API
	Server *http.Server
}

func New(cfg config.Config, registerRoutes httpapi.RouteRegistrar) *App {
	router, api := httpapi.NewRouter()
	if registerRoutes != nil {
		registerRoutes(api)
	}

	return &App{
		Config: cfg,
		API:    api,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.HTTPServer.Port),
			Handler: router,
		},
	}
}
