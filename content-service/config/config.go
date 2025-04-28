package config

import (
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server    Server `envPrefix:"SERVER_"`
		S3Storage S3Storage
	}

	Server struct {
		GRPCServer GRPCServer
		HTTPServer HTTPServer `envPrefix:"HTTP_"`
	}

	GRPCServer struct {
		Port                  int32         `env:"GRPC_PORT,notEmpty"`
		MaxRecvMsgSizeMiB     int           `env:"GRPC_MAX_MESSAGE_SIZE_MIB" envDefault:"12"`
		MaxConnectionAge      time.Duration `env:"GRPC_MAX_CONNECTION_AGE" envDefault:"30s"`
		MaxConnectionAgeGrace time.Duration `env:"GRPC_MAX_CONNECTION_AGE_GRACE" envDefault:"10s"`
	}

	HTTPServer struct {
		Port         int           `env:"HTTP_PORT" envDefault:"8080"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"30s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"30s"`
		Mode         string        `env:"GIN_MODE" envDefault:"release"` // release, debug, test
	}

	S3Storage struct {
		ConnStr       string `env:"S3_CONN_STR" envDefault:"http://127.0.0.1:4400"`
		DataDirectory string `env:"S3_DATA_DIRECTORY" envDefault:"./data"`
	}

	Nats struct{}
)

func New() (*Config, error) {
	var cfg Config
	opts := env.Options{Environment: nil, TagName: "env", Prefix: ""}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(&cfg, opts)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
