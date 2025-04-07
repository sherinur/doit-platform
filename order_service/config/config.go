package config

import (
	"time"

	"github.com/caarlos0/env"
)

type (
	Config struct {
		Server   Server
		Postgres Postgres
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

	Postgres struct {
		ConnStr string `env:"CONN_STR" envDefault:"user=admin password=root dbname=micro sslmode=disable"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
