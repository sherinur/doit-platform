package config

import (
	"time"

	postgresconn "github.com/sherinur/doit-platform/user-service/pkg/postgres"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Postgres  postgresconn.PostgresConfig
		Server    Server
		ZapLogger ZapLogger
		Jwt       Jwt

		Version string `env:"VERSION"`
	}

	Server struct {
		HTTPServer HTTPServer
		GRPCServer GRPCServer
	}

	HTTPServer struct {
		Port           int           `env:"HTTP_PORT,required"`
		ReadTimeout    time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"30s"`
		WriteTimeout   time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"30s"`
		IdleTimeout    time.Duration `env:"HTTP_IDLE_TIMEOUT" envDefault:"60s"`
		MaxHeaderBytes int           `env:"HTTP_MAX_HEADER_BYTES" envDefault:"1048576"` // 1 MB
		TrustedProxies []string      `env:"HTTP_TRUSTED_PROXIES" envSeparator:","`
		Mode           string        `env:"GIN_MODE" envDefault:"release"` // Can be: release, debug, test
	}

	GRPCServer struct {
		Port                  int32         `env:"GRPC_PORT,notEmpty"`
		MaxRecvMsgSizeMiB     int           `env:"GRPC_MAX_MESSAGE_SIZE_MIB" envDefault:"12"`
		MaxConnectionAge      time.Duration `env:"GRPC_MAX_CONNECTION_AGE" envDefault:"30s"`
		MaxConnectionAgeGrace time.Duration `env:"GRPC_MAX_CONNECTION_AGE_GRACE" envDefault:"10s"`
	}

	ZapLogger struct {
		Directory string `env:"ZAP_LOGGING_DIRECTORY" envDefault:"./logs"`
		Mode      string `env:"ZAP_LOGGING_MODE" envDefault:"debug"` // release, debug, test
	}

	Jwt struct {
		JwtAccessSecret      string `env:"JWT_ACCESS_SECRET"`
		JwtRefreshSecret     string `env:"JWT_REFRESH_SECRET"`
		JwtAccessExpiration  int    `env:"JWT_ACCESS_EXPIRATION"`
		JwtRefreshExpiration int    `env:"JWT_REFRESH_EXPIRATION"`
	}
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
