package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sherinur/doit-platform/user-service/config"
	cache "github.com/sherinur/doit-platform/user-service/internal/adapter/cahce"
	grpcserver "github.com/sherinur/doit-platform/user-service/internal/adapter/controller/grpc/server"
	"github.com/sherinur/doit-platform/user-service/internal/adapter/nats/producer"
	repo "github.com/sherinur/doit-platform/user-service/internal/adapter/repo/postgres"
	"github.com/sherinur/doit-platform/user-service/internal/usecase"
	natsconn "github.com/sherinur/doit-platform/user-service/pkg/nats"
	postgresconn "github.com/sherinur/doit-platform/user-service/pkg/postgres"
	"github.com/sherinur/doit-platform/user-service/pkg/security"
	"go.uber.org/zap"
)

const serviceName = "user-service"

type App struct {
	cfg *config.Config
	log *zap.Logger

	grpcServer *grpcserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	log.Printf("Starting %v service...", serviceName)

	// Connect to PostgreSQL
	log.Println("Connecting to PostgreSQL...")
	db, err := postgresconn.NewPostgreConnection(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	log.Println("Connected to PostgreSQL successfully.")

	// Connect to NATS
	log.Println("connecting to NATS", "hosts", strings.Join(cfg.Nats.Hosts, ","))
	natsClient, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats.NewClient: %w", err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	UserProducer := producer.NewUserProducer(natsClient)

	// Initialize Redis
	redisCache, err := cache.NewRedisCache(cfg.Redis.RedisAddr, cfg.Redis.RedisPass, cfg.Redis.RedisDB)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize redis cache: %w", err)
	}
	log.Println("Connecting to redis cahce")
	UserCache := cache.NewUserCache(redisCache)
	SessionCahce := cache.NewSessionCache(redisCache)

	// Initialize Repositories
	userRepo := repo.NewUserRepo(db)
	jwtManager := security.NewJWTManager(cfg.Jwt.JwtAccessSecret, cfg.Jwt.JwtRefreshSecret, cfg.Jwt.JwtAccessExpiration, cfg.Jwt.JwtRefreshExpiration)
	passwordManager := security.NewPasswordManager()

	// Initialize UseCases
	userUsecase := usecase.NewUserUsecase(userRepo, UserCache, SessionCahce, UserProducer, jwtManager, passwordManager)

	// Initialize HTTP Server
	grpcServer := grpcserver.New(cfg.Server, userUsecase, cfg.Jwt, logger)

	app := &App{
		grpcServer: grpcServer,
		log:        logger,
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

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	// Start the GRPC server
	a.grpcServer.Run(ctx, errCh)
	a.log.Info(fmt.Sprintf("server %v started", serviceName))

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
		a.log.Info(fmt.Sprintf("Received signal: %v. Running graceful shutdown...", s))
		a.Close(ctx)
		a.log.Info("Graceful shutdown completed!")
	}

	return nil
}
