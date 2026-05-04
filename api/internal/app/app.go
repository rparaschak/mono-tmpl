package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/internal/dependencies"
	"github.com/rparaschak/mono-tmpl/api/pkg/config"
	"github.com/rparaschak/mono-tmpl/api/pkg/httpapi"
	"github.com/rparaschak/mono-tmpl/api/pkg/mcpapi"
)

type App struct {
	Config       config.Config
	Deps         dependencies.Dependencies
	Modules      *Modules
	API          huma.API
	Server       *http.Server
	shutdownWait time.Duration
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	deps, err := dependencies.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	router, api := httpapi.NewRouter()

	modules := NewModules(deps)
	modules.RegisterHTTP(api)

	mcpServer := mcpapi.NewServer("Monorepo Template MCP", "1.0.0", modules.RegisterMCP)
	mcpapi.Mount(router, "/mcp", mcpServer)

	return &App{
		Config:  cfg,
		Deps:    deps,
		Modules: modules,
		API:     api,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.HTTPServer.Port),
			Handler: router,
		},
		shutdownWait: 10 * time.Second,
	}, nil
}

func (a *App) Handler() http.Handler {
	return a.Server.Handler
}

func (a *App) Run(ctx context.Context) error {
	serverErrors := make(chan error, 1)

	go func() {
		slog.Info("starting app", "addr", a.Server.Addr, "env", a.Config.HTTPServer.Env)
		if err := a.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
			return
		}
		serverErrors <- nil
	}()

	select {
	case err := <-serverErrors:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), a.shutdownWait)
		defer cancel()
		return a.Shutdown(shutdownCtx)
	}
}

func (a *App) Shutdown(ctx context.Context) error {
	var serverErr error
	if a.Server != nil {
		serverErr = a.Server.Shutdown(ctx)
	}

	return errors.Join(serverErr, a.Deps.Close())
}
