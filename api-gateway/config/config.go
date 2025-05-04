package config

import (
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server    Server
		ZapLogger ZapLogger
		Telemetry Telemetry
		GRPC      GRPC
		Version   string `env:"VERSION"`
	}

	Server struct {
		HTTPServer HTTPServer
	}

	HTTPServer struct {
		Port           int           `env:"HTTP_PORT,required"`
		ReadTimeout    time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"30s"`
		WriteTimeout   time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"30s"`
		IdleTimeout    time.Duration `env:"HTTP_IDLE_TIMEOUT" envDefault:"60s"`
		MaxHeaderBytes int           `env:"HTTP_MAX_HEADER_BYTES" envDefault:"1048576"` // 1 MB
		TrustedProxies []string      `env:"HTTP_TRUSTED_PROXIES" envSeparator:","`
		Mode           string        `env:"GIN_MODE" envDefault:"release"` // release, debug, test
	}

	GRPC struct {
		GRPCClient GRPCClient
	}

	GRPCClient struct {
		ContentServiceURL string `env:"GRPC_CONTENT_SERVICE_URL,required"`
	}

	ZapLogger struct {
		Directory string `env:"ZAP_LOGGING_DIRECTORY" envDefault:"./logs"`
		Mode      string `env:"ZAP_LOGGING_MODE" envDefault:"debug"` // release, debug, test
	}

	Telemetry struct {
		Mode                 string `env:"OTEL_MODE" envDefault:"debug"` // release, debug, test
		ExporterOTLPEndpoint string `env:"OTEL_EXPORTER_OTLP_ENDPOINT" envDefault:"http://localhost:4318"`
		ExporterOTLPInsecure bool   `env:"OTEL_EXPORTER_OTLP_INSECURE" envDefault:"true"`
		ExporterPromPort     int    `env:"OTEL_EXPORTER_PROM_PORT" envDefault:"3003"`
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
