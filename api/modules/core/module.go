package core

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/internal/dependencies"
	"github.com/rparaschak/mono-tmpl/api/modules/core/handlers"
)

type Module struct {
	Deps     dependencies.Dependencies
	Handlers *handlers.Handlers
}

func New(deps dependencies.Dependencies) *Module {
	return &Module{
		Deps:     deps,
		Handlers: &handlers.Handlers{},
	}
}

func (m *Module) RegisterHTTP(api huma.API) {
	RegisterRoutes(api, m.Handlers)
}
