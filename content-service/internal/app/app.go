package app

import (
	"context"
	"log/slog"
	"os"

	"content-service/config"
	"content-service/internal/adapter/controller/http/server"
	"content-service/internal/adapter/repo/s3"
	"content-service/internal/usecase"
)

const serviceName = "content-service"

type App struct {
	cfg *config.Config
	log *slog.Logger

	httpServer *server.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	fileRepo, err := s3.NewFile(cfg.S3Storage.ConnStr)
	if err != nil {
		return nil, err
	}

	fileUsecase := usecase.NewFile(fileRepo)
	httpServer := server.New(*cfg, fileUsecase)

	app := &App{
		log:        logger,
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the application")
	return a.httpServer.Run()
}

func (a *App) Stop() error {
	return nil
}
