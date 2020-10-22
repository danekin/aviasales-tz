package config

import "github.com/caarlos0/env"

type Config struct {
	HTTPServerConfig
}

type HTTPServerConfig struct {
	Port int `env:"HTTP_SERVER_PORT" envDefault:"8080"`
}

func Parse() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(&cfg.HTTPServerConfig); err != nil {
		return nil, err
	}

	return cfg, nil
}
