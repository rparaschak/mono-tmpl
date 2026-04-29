package sample

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/handlers"
	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

type Module struct {
	UseCases *usecases.UseCase
	Handlers *handlers.Handlers
}

func New() *Module {
	useCases := &usecases.UseCase{}
	return &Module{
		UseCases: useCases,
		Handlers: &handlers.Handlers{
			UseCases: useCases,
		},
	}
}

func (m *Module) WithRestRouter(api huma.API) *Module {
	RegisterRoutes(api, m.Handlers)
	return m
}
