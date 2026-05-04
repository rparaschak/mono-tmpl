package dependencies

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
)

func NewDefault(ctx context.Context, cfg config.Config) (Dependencies, error) {
	deps, err := NewEnv(cfg)
	if err != nil {
		return Dependencies{}, err
	}

	return deps, nil
}
