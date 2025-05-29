package app

import (
	"context"

	"github.com/sherinur/doit-platform/social-service/config"
	grpcserver "github.com/sherinur/doit-platform/social-service/internal/adapter/grpc/server"
	"go.uber.org/zap"
)

const serviceName = "social-service"

type App struct {
	cfg *config.Config
	log *zap.Logger

	grpcServer *grpcserver.API
	// httpServer *server.API

	telemetry *Telemetry
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	// logger
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	// controllers ...

	// telemetry
	telemetry, err := InitTelemetry(ctx, cfg.Telemetry, logger)
	if err != nil {
		return nil, err
	}

	app := &App{
		log:       logger,
		telemetry: telemetry,
	}

	return app, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the service")
	return a.grpcServer.Run(context.Background())
}

func (a *App) Stop() error {
	return nil
}
