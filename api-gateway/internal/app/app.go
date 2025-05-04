package app

import (
	"context"
	"fmt"

	"github.com/sherinur/doit-platform/api-gateway/config"
	"github.com/sherinur/doit-platform/api-gateway/internal/adapter/grpc/file"
	"github.com/sherinur/doit-platform/api-gateway/internal/adapter/http/server"
	"github.com/sherinur/doit-platform/api-gateway/internal/usecase"
	"github.com/sherinur/doit-platform/api-gateway/pkg/grpcconn"
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
	log, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}
	log.Debug("Zap logger is initialized", zap.String("mode", cfg.ZapLogger.Mode), zap.String("directory", cfg.ZapLogger.Directory))

	fileServiceGRPCConn, err := grpcconn.New(cfg.GRPC.GRPCClient.ContentServiceURL)
	if err != nil {
		return nil, err
	}
	log.Debug("Connected to the file service grpc client", zap.String("url", cfg.GRPC.GRPCClient.ContentServiceURL))

	filePresenter := file.NewFile(filesvc.NewFileServiceClient(fileServiceGRPCConn))
	fileUsecase := usecase.NewFile(filePresenter)

	httpServer := server.New(cfg.Server.HTTPServer, log, fileUsecase)

	app := &App{
		log:        log,
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
