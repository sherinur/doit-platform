package app

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"api-gateway/config"
	"api-gateway/internal/adapter/http/service"
)

const serviceName = "inventory_service"

type App struct {
	api *service.API
	cfg *config.Config
	log *slog.Logger
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	db, err := sql.Open("postgres", "postgres://admin:root@localhost/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

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
