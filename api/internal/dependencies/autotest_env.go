package dependencies

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
)

func NewAutotest(ctx context.Context, cfg config.Config) (Dependencies, error) {
	cfg.Storage.Endpoint = "localhost:5104"

	deps, err := NewEnv(cfg)
	if err != nil {
		return Dependencies{}, err
	}

	return deps, nil
}
