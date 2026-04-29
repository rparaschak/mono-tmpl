package handlers

import (
	"context"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/modules/sample/usecases"
)

func TestGetSamplesHandler(t *testing.T) {
	response, err := (Handlers{UseCases: &usecases.UseCase{}}).GetSamplesHandler(context.Background(), &GetSamplesRequest{})
	if err != nil {
		t.Fatalf("GetSamplesHandler() error = %v", err)
	}

	want := []string{"sample1", "sample2", "sample3"}
	if len(response.Body.Samples) != len(want) {
		t.Fatalf("len(Samples) = %d, want %d", len(response.Body.Samples), len(want))
	}

	for i := range want {
		if response.Body.Samples[i] != want[i] {
			t.Fatalf("Samples[%d] = %q, want %q", i, response.Body.Samples[i], want[i])
		}
	}
}
