package bootstrap

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"

	"github.com/rparaschak/mono-tmpl/api/modules"
	"github.com/rparaschak/mono-tmpl/api/pkg/openapi"
)

func TestRouteRegistrarRegistersModuleRoutesWithoutRuntimeDependencies(t *testing.T) {
	document, err := openapi.Document(func(api huma.API) {
		RegisterRoutes(api, modules.GlobalDependencies{})
	})
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

	if path.Post == nil {
		t.Fatal("OpenAPI document does not include POST /samples")
	}
	if path.Post.OperationID != "create-sample" {
		t.Fatalf("POST /samples OperationID = %q, want %q", path.Post.OperationID, "create-sample")
	}
	if requestBodySchemaRef(t, path.Post) != "#/components/schemas/SampleInputDTO" {
		t.Fatalf("POST /samples request body schema = %q, want SampleInputDTO", requestBodySchemaRef(t, path.Post))
	}
	if responseSchemaRef(t, path.Post, "201") != "#/components/schemas/SampleDTO" {
		t.Fatalf("POST /samples 201 response schema = %q, want SampleDTO", responseSchemaRef(t, path.Post, "201"))
	}

	itemPath := document.Paths["/samples/{sampleId}"]
	if itemPath == nil {
		t.Fatal("OpenAPI document does not include /samples/{sampleId}")
	}

	if itemPath.Put == nil {
		t.Fatal("OpenAPI document does not include PUT /samples/{sampleId}")
	}
	if itemPath.Put.OperationID != "update-sample" {
		t.Fatalf("PUT /samples/{sampleId} OperationID = %q, want %q", itemPath.Put.OperationID, "update-sample")
	}
	assertSampleIdPathParam(t, itemPath.Put)
	if requestBodySchemaRef(t, itemPath.Put) != "#/components/schemas/SampleInputDTO" {
		t.Fatalf("PUT /samples/{sampleId} request body schema = %q, want SampleInputDTO", requestBodySchemaRef(t, itemPath.Put))
	}
	if responseSchemaRef(t, itemPath.Put, "200") != "#/components/schemas/SampleDTO" {
		t.Fatalf("PUT /samples/{sampleId} 200 response schema = %q, want SampleDTO", responseSchemaRef(t, itemPath.Put, "200"))
	}
	if itemPath.Put.Responses["404"] == nil {
		t.Fatal("OpenAPI document does not include PUT /samples/{sampleId} 404 response")
	}

	if itemPath.Delete == nil {
		t.Fatal("OpenAPI document does not include DELETE /samples/{sampleId}")
	}
	if itemPath.Delete.OperationID != "delete-sample" {
		t.Fatalf("DELETE /samples/{sampleId} OperationID = %q, want %q", itemPath.Delete.OperationID, "delete-sample")
	}
	assertSampleIdPathParam(t, itemPath.Delete)
	if itemPath.Delete.Responses["204"] == nil {
		t.Fatal("OpenAPI document does not include DELETE /samples/{sampleId} 204 response")
	}
	if itemPath.Delete.Responses["404"] == nil {
		t.Fatal("OpenAPI document does not include DELETE /samples/{sampleId} 404 response")
	}

	schemas := document.Components.Schemas.Map()

	sampleDTO := schemas["SampleDTO"]
	if sampleDTO == nil {
		t.Fatal("OpenAPI document does not include SampleDTO schema")
	}
	if sampleDTO.Properties["latitude"] == nil {
		t.Fatal("SampleDTO schema does not include latitude")
	}
	if sampleDTO.Properties["longitude"] == nil {
		t.Fatal("SampleDTO schema does not include longitude")
	}

	sampleInputDTO := schemas["SampleInputDTO"]
	if sampleInputDTO == nil {
		t.Fatal("OpenAPI document does not include SampleInputDTO schema")
	}
	for _, field := range []string{"name", "latitude", "longitude"} {
		if sampleInputDTO.Properties[field] == nil {
			t.Fatalf("SampleInputDTO schema does not include %s", field)
		}
	}
}

func requestBodySchemaRef(t *testing.T, operation *huma.Operation) string {
	t.Helper()

	if operation.RequestBody == nil {
		t.Fatal("operation does not include request body")
	}

	mediaType := operation.RequestBody.Content["application/json"]
	if mediaType == nil || mediaType.Schema == nil {
		t.Fatal("operation does not include application/json request body schema")
	}

	return mediaType.Schema.Ref
}

func responseSchemaRef(t *testing.T, operation *huma.Operation, status string) string {
	t.Helper()

	response := operation.Responses[status]
	if response == nil {
		t.Fatalf("operation does not include %s response", status)
	}

	mediaType := response.Content["application/json"]
	if mediaType == nil || mediaType.Schema == nil {
		t.Fatalf("operation %s response does not include application/json schema", status)
	}

	return mediaType.Schema.Ref
}

func assertSampleIdPathParam(t *testing.T, operation *huma.Operation) {
	t.Helper()

	for _, parameter := range operation.Parameters {
		if parameter.Name == "sampleId" && parameter.In == "path" {
			if !parameter.Required {
				t.Fatal("sampleId path parameter is not required")
			}
			if parameter.Schema == nil || parameter.Schema.Format != "uuid" {
				t.Fatalf("sampleId path parameter format = %v, want uuid", parameter.Schema)
			}
			return
		}
	}

	t.Fatal("operation does not include sampleId path parameter")
}
