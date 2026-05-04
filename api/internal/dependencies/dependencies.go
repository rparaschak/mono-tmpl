package dependencies

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
	"github.com/rparaschak/mono-tmpl/api/pkg/database"
	"github.com/rparaschak/mono-tmpl/api/pkg/storage"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config  config.Config
	DB      *gorm.DB
	Storage storage.Service
}

func New(ctx context.Context, cfg config.Config) (Dependencies, error) {
	switch cfg.HTTPServer.Env {
	case "autotest":
		return NewAutotest(ctx, cfg)
	case "local":
		return NewLocal(ctx, cfg)
	default:
		return NewDefault(ctx, cfg)
	}
}

func NewEnv(_ context.Context, cfg config.Config) (Dependencies, error) {
	db, err := database.New(cfg.Database)
	if err != nil {
		return Dependencies{}, err
	}

	return Dependencies{
		Config: cfg,
		DB:     db,
	}, nil
}

func (d Dependencies) Close() error {
	if d.DB == nil {
		return nil
	}

	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
