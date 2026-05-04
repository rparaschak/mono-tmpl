package storage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	client *minio.Client
	region string
}

func New(cfg Config, buckets []string) (*Storage, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
		Region: cfg.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("create minio client: %w", err)
	}

	storage := &Storage{
		client: client,
		region: cfg.Region,
	}
	for _, bucket := range buckets {
		if _, err := storage.EnsureBucket(context.Background(), EnsureBucketInput{Bucket: bucket}); err != nil {
			return nil, err
		}
	}

	return storage, nil
}
