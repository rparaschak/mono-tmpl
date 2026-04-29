package config

import "testing"

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
}

func TestLoadHTTPServerDefaults(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.HTTPServer.Env != "local" {
		t.Fatalf("HTTPServer.Env = %q, want %q", cfg.HTTPServer.Env, "local")
	}

	if cfg.HTTPServer.Port != 5001 {
		t.Fatalf("HTTPServer.Port = %d, want %d", cfg.HTTPServer.Port, 5001)
	}
}
