//go:build integration

package storage_test

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/rparaschak/mono-tmpl/api/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUploadFileUseCase(t *testing.T) {
	ctx := context.Background()
	cfg := testConfig()
	bucket := uniqueBucket("upload")
	fileStorage, err := storage.New(cfg, []string{bucket})
	require.NoError(t, err, "storage should initialize")

	key := "documents/uploaded.txt"
	content := "uploaded file content"
	output, err := fileStorage.UploadFile(ctx, storage.UploadFileInput{
		Bucket:      bucket,
		Key:         key,
		Reader:      strings.NewReader(content),
		Size:        int64(len(content)),
		ContentType: "text/plain",
	})
	require.NoError(t, err, "uploading a single file should succeed")

	assert.Equal(t, bucket, output.Bucket, "upload output should include bucket")
	assert.Equal(t, key, output.Key, "upload output should include key")
	assert.Equal(t, int64(len(content)), output.Size, "upload output should include size")
	assert.Equal(t, "text/plain", output.ContentType, "upload output should include content type")

	client := testMinIOClient(t, cfg)
	object, err := client.GetObject(ctx, bucket, key, minio.GetObjectOptions{})
	require.NoError(t, err, "uploaded object should be readable from minio")
	defer object.Close()

	downloaded, err := io.ReadAll(object)
	require.NoError(t, err, "uploaded object body should be readable")
	assert.Equal(t, content, string(downloaded), "uploaded object body should match input")
}
