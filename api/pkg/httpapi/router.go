package httpapi

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type RouteRegistrar func(huma.API)

func NewRouter() (*http.ServeMux, huma.API) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	api := humago.New(mux, huma.DefaultConfig("Monorepo Template API", "1.0.0"))

	return mux, api
}
