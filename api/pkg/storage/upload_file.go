package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

type UploadFileInput struct {
	Bucket      string
	Key         string
	Reader      io.Reader
	Size        int64
	ContentType string
}

type UploadFileResult struct {
	Bucket      string
	Key         string
	Size        int64
	ContentType string
}

func (s *Storage) UploadFile(ctx context.Context, input UploadFileInput) (*UploadFileResult, error) {
	info, err := s.client.PutObject(ctx, input.Bucket, input.Key, input.Reader, input.Size, minio.PutObjectOptions{
		ContentType: input.ContentType,
	})
	if err != nil {
		return nil, fmt.Errorf("upload file: %w", err)
	}

	return &UploadFileResult{
		Bucket:      info.Bucket,
		Key:         info.Key,
		Size:        info.Size,
		ContentType: input.ContentType,
	}, nil
}
