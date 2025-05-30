package app

import (
	"context"
	"fmt"

	"github.com/sherinur/doit-platform/social-service/config"
	grpcserver "github.com/sherinur/doit-platform/social-service/internal/adapter/grpc/server"
	"github.com/sherinur/doit-platform/social-service/internal/adapter/inmemory"
	"github.com/sherinur/doit-platform/social-service/internal/adapter/mongo"
	"github.com/sherinur/doit-platform/social-service/internal/usecase"
	mongocon "github.com/sherinur/doit-platform/social-service/pkg/mongo"
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

	// mongo
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	feedbackRepo := mongo.NewFeedback(mongoDB.Conn)
	feedbackCache := inmemory.NewFeedback()
	feedbackUsecase := usecase.NewFeedback(feedbackRepo, feedbackCache)

	// controllers ...
	grpcServer, err := grpcserver.New(*cfg, logger, feedbackUsecase)
	if err != nil {
		return nil, err
	}

	// telemetry
	telemetry, err := InitTelemetry(ctx, cfg.Telemetry, logger)
	if err != nil {
		return nil, err
	}

	app := &App{
		log:        logger,
		telemetry:  telemetry,
		grpcServer: grpcServer,
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
