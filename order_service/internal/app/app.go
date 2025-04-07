package app

import (
	"context"
	"log/slog"
	"os"

	"order_service/config"
	"order_service/internal/adapter/http/service"
)

const serviceName = "order_service"

type App struct {
	api *service.API
	cfg *config.Config
	log *slog.Logger
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	api := service.New(&cfg.Server)

	return &App{log: logger, api: api}, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the application")
	return a.api.Run()
}

func (a *App) Stop() error {
	return nil
}
