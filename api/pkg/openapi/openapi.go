package openapi

import (
	"encoding/json"

	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/pkg/httpapi"
)

func Document(registerRoutes httpapi.RouteRegistrar) (*huma.OpenAPI, error) {
	_, api := httpapi.NewRouter()
	if registerRoutes != nil {
		registerRoutes(api)
	}
	return api.OpenAPI(), nil
}

func JSON(registerRoutes httpapi.RouteRegistrar) ([]byte, error) {
	document, err := Document(registerRoutes)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(document, "", "  ")
}
