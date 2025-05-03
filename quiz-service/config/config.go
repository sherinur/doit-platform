package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/sherinur/doit-platform/quiz-service/pkg/mongo"
)

type (
	Config struct {
		Mongo     mongo.Config
		Server    Server
		ZapLogger ZapLogger
	}

	Server struct {
		HttpServer HTTPServer
		GRPCServer GRPCServer
	}

	HTTPServer struct {
		Mode string `env:"GIN_MODE" envDefault:"release"` // release, debug, test
		Port string `env:"HTTP_PORT"`
	}

	GRPCServer struct {
		Port int32 `env:"GRPC_PORT"`
	}

	ZapLogger struct {
		Directory string `env:"ZAP_LOGGING_DIRECTORY" envDefault:"./logs"`
		Mode      string `env:"ZAP_LOGGING_MODE" envDefault:"./logs"` // release, debug, test
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg.Mongo)
	err = env.Parse(&cfg.Server)
	return &cfg, err

}
