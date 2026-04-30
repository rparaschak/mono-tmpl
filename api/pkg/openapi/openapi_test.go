package openapi_test

import (
	"context"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/rparaschak/mono-tmpl/api/pkg/openapi"
)

type testRequest struct{}

type testResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}

func TestDocumentRegistersRoutes(t *testing.T) {
	registerRoutes := func(api huma.API) {
		huma.Register(api, huma.Operation{
			OperationID: "get-test",
			Method:      "GET",
			Path:        "/test",
			Summary:     "Get Test",
		}, func(ctx context.Context, input *testRequest) (*testResponse, error) {
			return &testResponse{}, nil
		})
	}

	document, err := openapi.Document(registerRoutes)
	if err != nil {
		t.Fatalf("Document() error = %v", err)
	}

	path := document.Paths["/test"]
	if path == nil || path.Get == nil {
		t.Fatal("OpenAPI document does not include GET /test")
	}

	if path.Get.OperationID != "get-test" {
		t.Fatalf("OperationID = %q, want %q", path.Get.OperationID, "get-test")
	}
}
