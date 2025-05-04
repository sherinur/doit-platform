package app

import (
	"context"
	"fmt"

	"github.com/sherinur/doit-platform/api-gateway/config"
	"github.com/sherinur/doit-platform/api-gateway/internal/adapter/grpc/file"
	"github.com/sherinur/doit-platform/api-gateway/internal/adapter/http/server"
	filesvc "github.com/sherinur/doit-platform/apis/gen/content-service/service/frontend/file/v1"
	"go.uber.org/zap"
)

const serviceName = "api-gateway"

type App struct {
	cfg *config.Config
	log *zap.Logger

	httpServer *server.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	// logger
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	filePresenter := file.NewFile(filesvc.NewFileServiceClient())

	httpServer := server.New(cfg.Server.HTTPServer, logger)

	app := &App{
		log:        logger,
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Run() error {
	a.log.Info(fmt.Sprintf("Starting the %s service", serviceName))
	return a.httpServer.Run()
}

func (a *App) Stop() error {
	return nil
}
