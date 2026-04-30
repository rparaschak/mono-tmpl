package wiring

import (
	"testing"

	"github.com/rparaschak/mono-tmpl/api/modules"
	"github.com/rparaschak/mono-tmpl/api/pkg/openapi"
)

func TestRouteRegistrarRegistersModuleRoutesWithoutRuntimeDependencies(t *testing.T) {
	document, err := openapi.Document(RouteRegistrar(modules.GlobalDependencies{}))
	if err != nil {
		t.Fatalf("Document() error = %v", err)
	}

	path := document.Paths["/samples"]
	if path == nil || path.Get == nil {
		t.Fatal("OpenAPI document does not include GET /samples")
	}

	if path.Get.OperationID != "get-samples" {
		t.Fatalf("OperationID = %q, want %q", path.Get.OperationID, "get-samples")
	}
}
