package modules

import (
	"context"

	"github.com/rparaschak/mono-tmpl/api/pkg/config"
	"github.com/rparaschak/mono-tmpl/api/pkg/database"
	"gorm.io/gorm"
)

type GlobalDependencies struct {
	Config config.Config
	DB     *gorm.DB
}

func NewDependencies(ctx context.Context, cfg config.Config) (GlobalDependencies, error) {
	dbConfig := cfg.Database
	dbConfig.Env = cfg.HTTPServer.Env

	return GlobalDependencies{
		Config: cfg,
		DB:     database.New(dbConfig),
	}, nil
}
