package app

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"honnef.co/go/tools/config"
)

const serviceName = "content-service"

type App struct {
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

	return &App{log: logger}, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the application")
	return nil
}

func (a *App) Stop() error {
	return nil
}
