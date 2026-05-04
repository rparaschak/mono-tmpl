package core

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/modules/core/handlers"
	"github.com/rparaschak/mono-tmpl/api/pkg/routing"
)

func RegisterRoutes(parentRouter huma.API, h *handlers.Handlers) {
	groups := routing.NewBuilder(parentRouter, "/core", "Core").Groups()

	huma.Register(groups.Public, huma.Operation{
		OperationID:   "ping",
		Method:        http.MethodGet,
		Path:          "/ping",
		Summary:       "Ping Core",
		DefaultStatus: http.StatusNoContent,
	}, h.PingHandler)
}
