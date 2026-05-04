//go:build integration

package storage_test

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadFileUseCase(t *testing.T) {
	ctx := context.Background()
	cfg := testConfig()
	bucket := uniqueBucket("download")
	fileStorage, err := storage.New(cfg, []string{bucket})
	require.NoError(t, err, "storage should initialize")

	key := "documents/downloaded.txt"
	content := "downloaded file content"
	_, err = fileStorage.UploadFile(ctx, storage.UploadFileInput{
		Bucket:      bucket,
		Key:         key,
		Reader:      strings.NewReader(content),
		Size:        int64(len(content)),
		ContentType: "text/plain",
	})
	require.NoError(t, err, "test file upload should succeed")

	output, err := fileStorage.DownloadFile(ctx, storage.DownloadFileInput{
		Bucket: bucket,
		Key:    key,
	})
	require.NoError(t, err, "downloading a single file should succeed")
	defer output.Reader.Close()

	downloaded, err := io.ReadAll(output.Reader)
	require.NoError(t, err, "downloaded object body should be readable")

	assert.Equal(t, bucket, output.Bucket, "download output should include bucket")
	assert.Equal(t, key, output.Key, "download output should include key")
	assert.Equal(t, int64(len(content)), output.Size, "download output should include size")
	assert.Equal(t, "text/plain", output.ContentType, "download output should include content type")
	assert.Equal(t, content, string(downloaded), "downloaded object body should match uploaded content")
}
