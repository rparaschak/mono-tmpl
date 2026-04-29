package sample

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/handlers"
	"github.com/rparaschak/mono-tmpl/api/pkg/routing"
)

func RegisterRoutes(parentRouter huma.API, h *handlers.Handlers) {
	groups := routing.NewBuilder(parentRouter, "/samples", "Samples").Groups()

	routing.GET(groups.Public, "", "Get Samples", h.GetSamplesHandler)
}
