package storage

import "context"

type Service interface {
	Upload(ctx context.Context) error
}

type MockService struct{}

func NewMockService() *MockService {
	return &MockService{}
}

func (s *MockService) Upload(_ context.Context) error {
	return nil
}
