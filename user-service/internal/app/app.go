package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	userrepo "user-services/internal/adapter/repo/postgres"
	"user-services/internal/usecase"
	"user-services/pkg/postgres"

	"github.com/sherinur/doit-platform/user-service/config"
	httpserver "github.com/sherinur/doit-platform/user-service/internal/adapter/controller/http/server"
)

const serviceName = "user-service"

type App struct {
	httpServer *httpserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("Starting %v service...", serviceName)

	// Connect to PostgreSQL
	log.Println("Connecting to PostgreSQL...")
	db, err := postgres.NewPostgreConnection(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	log.Println("Connected to PostgreSQL successfully.")

	// Initialize Repositories
	userRepo := userrepo.NewUserRepo(db)

	// Initialize UseCases
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Initialize HTTP Server
	httpServer := httpserver.New(cfg.Server, userUsecase)

	app := &App{
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Close(ctx context.Context) {
	// Stop the HTTP server
	err := a.httpServer.Stop()
	if err != nil {
		log.Println("Failed to shutdown HTTP server:", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()

	// Start the HTTP server
	a.httpServer.Run(errCh)
	log.Printf("Service %v started", serviceName)

	// Wait for termination signals
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Printf("Received signal: %v. Running graceful shutdown...", s)
		a.Close(ctx)
		log.Println("Graceful shutdown completed!")
	}

	return nil
}
