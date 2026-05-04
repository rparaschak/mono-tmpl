package dependencies

import (
	"context"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/pkg/appenv"
	"github.com/rparaschak/mono-tmpl/api/pkg/config"
)

func TestNewReturnsDatabaseErrors(t *testing.T) {
	cfg := config.Config{
		HTTPServer: config.HTTPServerConfig{
			Env: appenv.Autotest,
		},
	}

	if _, err := New(context.Background(), cfg); err == nil {
		t.Fatal("New() error = nil, want database initialization error")
	}
}

func TestCloseAllowsEmptyDependencies(t *testing.T) {
	if err := (Dependencies{}).Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}
