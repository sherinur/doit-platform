package app

import (
	"context"
	"fmt"
	"log"

	"github.com/sherinur/doit-platform/course-service/config"
	grpcserver "github.com/sherinur/doit-platform/course-service/internal/adapters/controller/grpc/server"
	mongoRepo "github.com/sherinur/doit-platform/course-service/internal/adapters/repo/mongo"
	"github.com/sherinur/doit-platform/course-service/internal/usecase"
	mongocon "github.com/sherinur/doit-platform/course-service/pkg/mongo"
	"go.uber.org/zap"
)

const serviceName = "content-service"

type App struct {
	cfg *config.Config
	log *zap.Logger

	grpcServer *grpcserver.API

	telemetry *Telemetry
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	// logger
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	// Initialize database connection here
	log.Println("connecting to mongo", "database", cfg.Mongo.Database)
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}
	courseRepo := mongoRepo.NewCourseRepository(mongoDB.Conn)
	courseUsecase := usecase.NewCourseUsecase(courseRepo)
	// controllers
	grpcServer := grpcserver.New(*cfg, courseUsecase, logger)

	// telemetry
	telemetry, err := InitTelemetry(ctx, cfg.Telemetry, logger)
	if err != nil {
		return nil, err
	}

	app := &App{
		log:        logger,
		grpcServer: grpcServer,
		telemetry:  telemetry,
	}

	return app, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the service")
	return a.grpcServer.Run(context.Background())
}

func (a *App) Stop() error {
	return nil
}
