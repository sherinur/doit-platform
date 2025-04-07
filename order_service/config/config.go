package config

import (
	"time"

	"github.com/caarlos0/env"
)

type (
	Config struct {
		Server Server
	}

	Server struct {
		HTTPServer HTTPServer
	}

	HTTPServer struct {
		Port         int           `env:"HTTP_PORT" envDefault:"8080"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"30s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"30s"`
		Mode         string        `env:"GIN_MODE" envDefault:"release"` // release, debug, test
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
