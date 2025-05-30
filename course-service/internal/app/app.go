package app

import (
	"context"

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

	// TODO : Initialize database connection here
	mongoClient := mongocon.Connect(cfg.Mongo.URI)
	courseRepo := mongoRepo.NewCourseRepository(mongoCliet.)
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
