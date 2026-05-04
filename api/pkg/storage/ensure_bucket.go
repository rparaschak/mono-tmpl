package storage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

type EnsureBucketInput struct {
	Bucket string
}

type EnsureBucketResult struct {
	Bucket  string
	Created bool
}

func (s *Storage) EnsureBucket(ctx context.Context, input EnsureBucketInput) (*EnsureBucketResult, error) {
	exists, err := s.client.BucketExists(ctx, input.Bucket)
	if err != nil {
		return nil, fmt.Errorf("check bucket %q: %w", input.Bucket, err)
	}
	if exists {
		return &EnsureBucketResult{Bucket: input.Bucket}, nil
	}

	if err := s.client.MakeBucket(ctx, input.Bucket, minio.MakeBucketOptions{Region: s.region}); err != nil {
		return nil, fmt.Errorf("create bucket %q: %w", input.Bucket, err)
	}

	return &EnsureBucketResult{Bucket: input.Bucket, Created: true}, nil
}
