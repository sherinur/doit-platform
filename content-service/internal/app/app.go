package app

import (
	"content-service/config"
	"context"
	"log/slog"
	"os"
)

const serviceName = "content-service"

type App struct {
	cfg *config.Config
	log *slog.Logger
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return &App{log: logger}, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the application")
	return nil
}

func (a *App) Stop() error {
	return nil
}
