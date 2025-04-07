package app

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"order_service/config"
	"order_service/internal/adapter/http/service"
	"order_service/internal/adapter/http/service/handler"
	"order_service/internal/adapter/postgres"
	"order_service/internal/usecase"

	_ "github.com/lib/pq"
)

const serviceName = "order_service"

type App struct {
	api *service.API
	cfg *config.Config
	log *slog.Logger
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	db, err := sql.Open("postgres", "postgres://admin:root@localhost/micro?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	orderRepo := postgres.NewOrder(db)
	orderUsecase := usecase.NewOrder(orderRepo)
	orderHandler := handler.NewOrder(orderUsecase)

	api := service.New(&cfg.Server, orderHandler)

	return &App{log: logger, api: api}, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the application")
	return a.api.Run()
}

func (a *App) Stop() error {
	return nil
}
