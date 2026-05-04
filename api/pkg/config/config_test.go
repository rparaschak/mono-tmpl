package config

import (
	"testing"

	"github.com/rparaschak/mono-tmpl/api/pkg/appenv"
)

func TestLoadHTTPServerFromEnv(t *testing.T) {
	t.Setenv("APP_ENV", "test")
	t.Setenv("APP_PORT", "7000")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.HTTPServer.Env != "test" {
		t.Fatalf("HTTPServer.Env = %q, want %q", cfg.HTTPServer.Env, "test")
	}

	if cfg.HTTPServer.Port != 7000 {
		t.Fatalf("HTTPServer.Port = %d, want %d", cfg.HTTPServer.Port, 7000)
	}

	if cfg.Database.Env != "test" {
		t.Fatalf("Database.Env = %q, want %q", cfg.Database.Env, "test")
	}
}

func TestLoadHTTPServerDefaults(t *testing.T) {
	t.Setenv("APP_ENV", "")
	t.Setenv("APP_PORT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.HTTPServer.Env != appenv.Local {
		t.Fatalf("HTTPServer.Env = %q, want %q", cfg.HTTPServer.Env, appenv.Local)
	}

	if cfg.HTTPServer.Port != 5001 {
		t.Fatalf("HTTPServer.Port = %d, want %d", cfg.HTTPServer.Port, 5001)
	}
}
