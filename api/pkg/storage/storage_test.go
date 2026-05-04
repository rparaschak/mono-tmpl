package storage

import (
	"context"
	"testing"
)

func TestMockServiceUpload(t *testing.T) {
	service := NewMockService()

	if err := service.Upload(context.Background()); err != nil {
		t.Fatalf("Upload() error = %v", err)
	}
}
