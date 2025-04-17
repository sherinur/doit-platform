package app

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"inventory_service/config"
	"inventory_service/internal/adapter/http/service"
	"inventory_service/internal/adapter/http/service/handler"
	"inventory_service/internal/adapter/postgres"
	"inventory_service/internal/usecase"

	_ "github.com/lib/pq"
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

	productRepo := postgres.NewProduct(db)
	productUsecase := usecase.NewProduct(productRepo)
	productHandler := handler.NewProduct(productUsecase)

	discountRepo := postgres.NewDiscount(db)
	discountUsecase := usecase.NewDiscount(discountRepo)
	discountHandler := handler.NewDiscount(discountUsecase)

	api := service.New(&cfg.Server, productHandler, discountHandler)

	return &App{log: logger, api: api}, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the application")
	return a.api.Run()
}

func (a *App) Stop() error {
	return nil
}
