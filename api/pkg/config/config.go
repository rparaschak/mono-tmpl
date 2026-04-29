package config

import "github.com/caarlos0/env/v11"

type Config struct {
	HTTPServer HTTPServerConfig
}

type HTTPServerConfig struct {
	Env  string `env:"APP_ENV"  envDefault:"local"`
	Port int    `env:"APP_PORT" envDefault:"5001"`
}

func Load() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
