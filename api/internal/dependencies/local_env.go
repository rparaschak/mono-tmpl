package dependencies

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
	"github.com/rparaschak/mono-tmpl/api/pkg/storage"
)

func NewLocal(ctx context.Context, cfg config.Config) (Dependencies, error) {
	deps, err := NewEnv(ctx, cfg)
	if err != nil {
		return Dependencies{}, err
	}

	deps.Storage = storage.NewMockService()
	return deps, nil
}
