package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	URL                string        `env:"DATABASE_URL"         envDefault:"postgresql://supabase_admin:docker@localhost:5002/mono-tmpl?search_path=public"`
	MaxConnections     int           `env:"DB_MAX_CONNECTIONS"   envDefault:"20"`
	MaxIdleConnections int           `env:"DB_MAX_IDLE"          envDefault:"10"`
	MaxLifetime        time.Duration `env:"DB_MAX_LIFETIME"      envDefault:"15m"`
	SlowThreshold      time.Duration `env:"DB_SLOW_THRESHOLD"    envDefault:"1ms"`
	Env                string        `env:"APP_ENV"              envDefault:"local"`
}

func New(config Config) *gorm.DB {
	var gormLogger logger.Interface
	switch config.Env {
	case "autotest":
		gormLogger = logger.Default.LogMode(logger.Silent)
	case "local":
		gormLogger = logger.Default.LogMode(logger.Info)
	default:
		gormLogger = logger.Default.LogMode(logger.Warn)
	}

	db, err := gorm.Open(postgres.Open(config.URL), &gorm.Config{
		Logger:         gormLogger,
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.MaxConnections)
	sqlDB.SetConnMaxLifetime(config.MaxLifetime)

	return db
}
