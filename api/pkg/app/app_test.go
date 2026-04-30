package app

import (
	"net/http"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
)

func TestNewUsesCodeProvidedConfig(t *testing.T) {
	cfg := config.Config{
		HTTPServer: config.HTTPServerConfig{
			Env:  "test",
			Port: 8080,
		},
	}

	application := New(cfg, nil)

	if application.Config != cfg {
		t.Fatalf("Config = %#v, want %#v", application.Config, cfg)
	}

	if application.Server.Addr != ":8080" {
		t.Fatalf("Server.Addr = %q, want %q", application.Server.Addr, ":8080")
	}

	if application.Server.Handler == nil {
		t.Fatal("Server.Handler is nil")
	}

	if _, ok := application.Server.Handler.(*http.ServeMux); !ok {
		t.Fatalf("Server.Handler type = %T, want *http.ServeMux", application.Server.Handler)
	}
}
