package app

import (
	"context"

	"github.com/sherinur/doit-platform/content-service/config"
	grpcserver "github.com/sherinur/doit-platform/content-service/internal/adapter/controller/grpc/server"
	"github.com/sherinur/doit-platform/content-service/internal/adapter/controller/http/server"
	"github.com/sherinur/doit-platform/content-service/internal/adapter/repo/s3"
	"github.com/sherinur/doit-platform/content-service/internal/usecase"
	"go.uber.org/zap"
)

const serviceName = "content-service"

type App struct {
	cfg *config.Config
	log *zap.Logger

	grpcServer *grpcserver.API
	httpServer *server.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	fileRepo, err := s3.NewFile(cfg.S3Storage.ConnStr)
	if err != nil {
		return nil, err
	}

	fileUsecase := usecase.NewFile(fileRepo)

	// controllers
	httpServer := server.New(*cfg, fileUsecase)
	grpcServer := grpcserver.New(*cfg, fileUsecase, logger)

	app := &App{
		log:        logger,
		httpServer: httpServer,
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
