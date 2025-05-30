package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/sherinur/doit-platform/quiz-service/pkg/mongo"
	"github.com/sherinur/doit-platform/quiz-service/pkg/redis"
)

type (
	Config struct {
		Mongo     mongo.Config
		Redis     redis.Config
		Server    Server
		Nats      Nats
		ZapLogger ZapLogger
		Telemetry Telemetry
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

	Nats struct {
		Hosts        []string `env:"NATS_HOSTS,notEmpty" envSeparator:","`
		NKey         string   `env:"NATS_NKEY,notEmpty"`
		IsTest       bool     `env:"NATS_IS_TEST,notEmpty" envDefault:"true"`
		NatsSubjects NatsSubjects
	}

	NatsSubjects struct {
		ClientEventSubject string `env:"NATS_COURSE_EVENT_SUBJECT,notEmpty"`
	}

	ZapLogger struct {
		Directory string `env:"ZAP_LOGGING_DIRECTORY" envDefault:"./logs"`
		Mode      string `env:"ZAP_LOGGING_MODE" envDefault:"release"` // release, debug, test
	}

	Telemetry struct {
		Mode                 string `env:"OTEL_MODE" envDefault:"debug"` // release, debug, test
		ExporterOTLPEndpoint string `env:"OTEL_EXPORTER_OTLP_ENDPOINT" envDefault:"http://localhost:4318"`
		ExporterOTLPInsecure bool   `env:"OTEL_EXPORTER_OTLP_INSECURE" envDefault:"true"`
		ExporterPromPort     int    `env:"OTEL_EXPORTER_PROM_PORT" envDefault:"3002"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg.Mongo)
	err = env.Parse(&cfg.Server)
	err = env.Parse(&cfg.ZapLogger)
	err = env.Parse(&cfg.Telemetry)
	err = env.Parse(&cfg.Nats)
	err = env.Parse(&cfg.Nats.NatsSubjects)

	return &cfg, err

}
