package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sherinur/doit-platform/user-service/config"
	grpcserver "github.com/sherinur/doit-platform/user-service/internal/adapter/controller/grpc/server"
	repo "github.com/sherinur/doit-platform/user-service/internal/adapter/repo/postgres"
	"github.com/sherinur/doit-platform/user-service/internal/usecase"
	postgresconn "github.com/sherinur/doit-platform/user-service/pkg/postgres"
	"github.com/sherinur/doit-platform/user-service/pkg/security"
)

const serviceName = "user-service"

type App struct {
	// httpServer *httpserver.API
	grpcServer *grpcserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("Starting %v service...", serviceName)

	// Connect to PostgreSQL
	log.Println("Connecting to PostgreSQL...")
	db, err := postgresconn.NewPostgreConnection(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	log.Println("Connected to PostgreSQL successfully.")

	// Initialize Repositories
	userRepo := repo.NewUserRepo(db)
	tokenRepo := repo.NewSessionRepo(db)
	jwtManager := security.NewJWTManager(cfg.Jwt.JwtAccessSecret, cfg.Jwt.JwtRefreshSecret, cfg.Jwt.JwtAccessExpiration, cfg.Jwt.JwtRefreshExpiration)
	passwordManager := security.NewPasswordManager()

	// Initialize UseCases
	userUsecase := usecase.NewUserUsecase(userRepo, tokenRepo, jwtManager, passwordManager)

	// Initialize HTTP Server
	grpcServer := grpcserver.New(cfg.Server, userUsecase, cfg.Jwt)

	app := &App{
		grpcServer: grpcServer,
	}

	return app, nil
}

func (a *App) Close(ctx context.Context) {
	err := a.grpcServer.Stop(context.Background())
	if err != nil {
		log.Println("failed to shutdown server", err)
	}

	// Stop the HTTP server
	// err = a.httpServer.Stop()
	// if err != nil {
	// 	log.Println("Failed to shutdown HTTP server:", err)
	// }
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()

	// Start the GRPC server
	a.grpcServer.Run(ctx, errCh)
	log.Println(fmt.Sprintf("server %v started", serviceName))

	// Start the HTTP server
	// a.httpServer.Run(errCh)
	// log.Printf("Service %v started", serviceName)

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
