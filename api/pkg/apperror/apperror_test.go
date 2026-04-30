package apperror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/pkg/apperror"
)

func TestAppErrorImplementsStatusError(t *testing.T) {
	err := apperror.New(http.StatusNotFound, "sample_not_found", "sample not found")

	if err.GetStatus() != http.StatusNotFound {
		t.Fatalf("GetStatus() = %d, want %d", err.GetStatus(), http.StatusNotFound)
	}
	if err.Error() != "sample not found" {
		t.Fatalf("Error() = %q, want %q", err.Error(), "sample not found")
	}
}

func TestWithDetailsCopiesBaseError(t *testing.T) {
	base := apperror.New(http.StatusNotFound, "sample_not_found", "sample not found")

	detailed := base.WithDetails(map[string]any{"sampleId": "missing-id"})

	if base.Details != nil {
		t.Fatal("base error details should remain nil")
	}
	if detailed.Details == nil {
		t.Fatal("detailed error should include details")
	}
	if detailed.GetStatus() != base.GetStatus() {
		t.Fatalf("detailed status = %d, want %d", detailed.GetStatus(), base.GetStatus())
	}
	if !errors.Is(detailed, base) {
		t.Fatal("detailed error should match base error")
	}
}

func TestErrorIsOneOf(t *testing.T) {
	notFound := apperror.New(http.StatusNotFound, "sample_not_found", "sample not found")
	conflict := apperror.New(http.StatusConflict, "sample_conflict", "sample conflict")

	if !apperror.ErrorIsOneOf(notFound.WithDetails("missing-id"), conflict, notFound) {
		t.Fatal("expected detailed not-found error to match one candidate")
	}
}
