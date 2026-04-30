package sample

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/handlers"
	"github.com/rparaschak/mono-tmpl/api/pkg/routing"
)

func RegisterRoutes(parentRouter huma.API, h *handlers.Handlers) {
	groups := routing.NewBuilder(parentRouter, "/samples", "Samples").Groups()

	routing.GET(groups.Public, "", "Get Samples", h.GetSamplesHandler)
	routing.POST(groups.Public, "", "Create Sample", h.CreateSampleHandler, routing.WithDefaultStatus(http.StatusCreated))
	routing.PUT(groups.Public, "/{sampleId}", "Update Sample", h.UpdateSampleHandler, routing.WithErrors(http.StatusNotFound))
	routing.DELETE(groups.Public, "/{sampleId}", "Delete Sample", h.DeleteSampleHandler, routing.WithErrors(http.StatusNotFound))
	routing.GET(groups.Public, "/nearby", "Get Nearby Samples", h.GetNearbySamplesHandler)
}
