//go:build integration

package handlers_test

import (
	"net/http"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/internal/testsupport/integration"
)

func TestPingAPI(t *testing.T) {
	env := integration.NewAutotestEnv(t)

	t.Run("GET /core/ping", func(t *testing.T) {
		t.Run("returns no content", func(t *testing.T) {
			env.Expect.GET("/core/ping").
				Expect().
				Status(http.StatusNoContent)
		})
	})

	t.Run("GET /health is not registered", func(t *testing.T) {
		t.Run("returns not found", func(t *testing.T) {
			env.Expect.GET("/health").
				Expect().
				Status(http.StatusNotFound)
		})
	})
}
