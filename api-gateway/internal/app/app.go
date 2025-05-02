package app

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/config"
	"go.uber.org/zap"
)

const serviceName = "api-gateway"

type App struct {
	cfg *config.Config
	log *zap.Logger

	// telemetry *Telemetry
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	// logger
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}
	// // telemetry
	// telemetry, err := InitTelemetry(ctx, cfg.Telemetry, logger)
	// if err != nil {
	// 	return nil, err
	// }

	app := &App{
		log: logger,
		// httpServer: httpServer,
		// grpcServer: grpcServer,
		// telemetry:  telemetry,
	}

	return app, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the service")
	// return a.grpcServer.Run(context.Background())
	return nil
}

func (a *App) Stop() error {
	return nil
}
