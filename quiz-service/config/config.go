package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/sherinur/doit-platform/quiz-service/pkg/mongo"
)

type (
	Config struct {
		Mongo  mongo.Config
		Server Server
	}

	Server struct {
		HttpServer HTTPServer
		GRPCServer GRPCServer
	}

	HTTPServer struct {
		Mode string `env:"GIN_MODE" envDefault:"release"`
		Port string `env:"HTTP_PORT"`
	}

	GRPCServer struct {
		Port int32 `env:"GRPC_PORT"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg.Mongo)
	err = env.Parse(&cfg.Server)
	return &cfg, err

}
