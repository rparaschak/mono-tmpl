package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

type DownloadFileInput struct {
	Bucket string
	Key    string
}

type DownloadFileResult struct {
	Bucket      string
	Key         string
	Reader      io.ReadCloser
	Size        int64
	ContentType string
}

func (s *Storage) DownloadFile(ctx context.Context, input DownloadFileInput) (*DownloadFileResult, error) {
	object, err := s.client.GetObject(ctx, input.Bucket, input.Key, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("download file: %w", err)
	}

	info, err := object.Stat()
	if err != nil {
		_ = object.Close()
		return nil, fmt.Errorf("stat downloaded file: %w", err)
	}

	return &DownloadFileResult{
		Bucket:      input.Bucket,
		Key:         input.Key,
		Reader:      object,
		Size:        info.Size,
		ContentType: info.ContentType,
	}, nil
}
