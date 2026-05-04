//go:build integration

package storage_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rparaschak/mono-tmpl/api/pkg/storage"
	"github.com/stretchr/testify/require"
)

func testConfig() storage.Config {
	return storage.Config{
		Endpoint:        "localhost:5104",
		AccessKeyID:     "minioadmin",
		SecretAccessKey: "minioadmin",
		UseSSL:          false,
		Region:          "us-east-1",
	}
}

func testMinIOClient(t *testing.T, cfg storage.Config) *minio.Client {
	t.Helper()

	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
		Region: cfg.Region,
	})
	require.NoError(t, err, "minio test client should initialize")

	return client
}

func uniqueBucket(prefix string) string {
	return prefix + "-" + uuid.NewString()
}
