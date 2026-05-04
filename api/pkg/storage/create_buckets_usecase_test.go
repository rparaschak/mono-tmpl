//go:build integration

package storage_test

import (
	"context"
	"testing"

	"github.com/rparaschak/mono-tmpl/api/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCreatesBucketsUseCase(t *testing.T) {
	ctx := context.Background()
	cfg := testConfig()
	bucketA := uniqueBucket("init-a")
	bucketB := uniqueBucket("init-b")

	_, err := storage.New(cfg, []string{bucketA, bucketB})
	require.NoError(t, err, "storage should initialize and create missing buckets")

	client := testMinIOClient(t, cfg)
	existsA, err := client.BucketExists(ctx, bucketA)
	require.NoError(t, err, "first bucket existence check should succeed")
	existsB, err := client.BucketExists(ctx, bucketB)
	require.NoError(t, err, "second bucket existence check should succeed")

	assert.True(t, existsA, "first configured bucket should exist")
	assert.True(t, existsB, "second configured bucket should exist")
}

func TestEnsureBucketUseCase(t *testing.T) {
	ctx := context.Background()
	cfg := testConfig()
	fileStorage, err := storage.New(cfg, nil)
	require.NoError(t, err, "storage should initialize without configured buckets")

	bucket := uniqueBucket("ensure")
	result, err := fileStorage.EnsureBucket(ctx, storage.EnsureBucketInput{Bucket: bucket})
	require.NoError(t, err, "ensuring a missing bucket should succeed")

	assert.Equal(t, bucket, result.Bucket, "ensure bucket result should include bucket")
	assert.True(t, result.Created, "ensure bucket result should report created bucket")

	secondResult, err := fileStorage.EnsureBucket(ctx, storage.EnsureBucketInput{Bucket: bucket})
	require.NoError(t, err, "ensuring an existing bucket should succeed")

	assert.Equal(t, bucket, secondResult.Bucket, "second ensure bucket result should include bucket")
	assert.False(t, secondResult.Created, "second ensure bucket result should report existing bucket")
}
