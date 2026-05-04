package config

import (
	"testing"

	"github.com/rparaschak/mono-tmpl/api/pkg/appenv"
)

func TestLoadHTTPServerFromEnv(t *testing.T) {
	t.Setenv("APP_ENV", "test")
	t.Setenv("APP_PORT", "7000")
	t.Setenv("STORAGE_ENDPOINT", "storage.test:9000")
	t.Setenv("STORAGE_ACCESS_KEY_ID", "test-access-key")
	t.Setenv("STORAGE_SECRET_ACCESS_KEY", "test-secret-key")
	t.Setenv("STORAGE_USE_SSL", "true")
	t.Setenv("STORAGE_REGION", "eu-central-1")
	t.Setenv("STORAGE_BUCKETS", "files,avatars")

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

	if cfg.Storage.Endpoint != "storage.test:9000" {
		t.Fatalf("Storage.Endpoint = %q, want %q", cfg.Storage.Endpoint, "storage.test:9000")
	}

	if cfg.Storage.AccessKeyID != "test-access-key" {
		t.Fatalf("Storage.AccessKeyID = %q, want %q", cfg.Storage.AccessKeyID, "test-access-key")
	}

	if cfg.Storage.SecretAccessKey != "test-secret-key" {
		t.Fatalf("Storage.SecretAccessKey = %q, want %q", cfg.Storage.SecretAccessKey, "test-secret-key")
	}

	if !cfg.Storage.UseSSL {
		t.Fatal("Storage.UseSSL = false, want true")
	}

	if cfg.Storage.Region != "eu-central-1" {
		t.Fatalf("Storage.Region = %q, want %q", cfg.Storage.Region, "eu-central-1")
	}

	if len(cfg.Storage.Buckets) != 2 || cfg.Storage.Buckets[0] != "files" || cfg.Storage.Buckets[1] != "avatars" {
		t.Fatalf("Storage.Buckets = %#v, want %#v", cfg.Storage.Buckets, []string{"files", "avatars"})
	}
}

func TestLoadHTTPServerDefaults(t *testing.T) {
	t.Setenv("APP_ENV", "")
	t.Setenv("APP_PORT", "")
	t.Setenv("STORAGE_ENDPOINT", "")
	t.Setenv("STORAGE_ACCESS_KEY_ID", "")
	t.Setenv("STORAGE_SECRET_ACCESS_KEY", "")
	t.Setenv("STORAGE_USE_SSL", "")
	t.Setenv("STORAGE_REGION", "")
	t.Setenv("STORAGE_BUCKETS", "")

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

	if cfg.Storage.Endpoint != "localhost:5004" {
		t.Fatalf("Storage.Endpoint = %q, want %q", cfg.Storage.Endpoint, "localhost:5004")
	}

	if len(cfg.Storage.Buckets) != 1 || cfg.Storage.Buckets[0] != "files" {
		t.Fatalf("Storage.Buckets = %#v, want %#v", cfg.Storage.Buckets, []string{"files"})
	}
}
