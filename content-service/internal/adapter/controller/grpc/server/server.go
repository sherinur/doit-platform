package grpcserver

import (
	"context"
	"fmt"
	"net"

	svc "github.com/sherinur/doit-platform/apis/gen/content-service/service/frontend/file/v1"
	"github.com/sherinur/doit-platform/content-service/config"
	"github.com/sherinur/doit-platform/content-service/internal/adapter/controller/grpc/server/frontend"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverIPAddress = "0.0.0.0:%d"

type API struct {
	server *grpc.Server
	cfg    config.GRPCServer
	addr   string

	log *zap.Logger

	fileUsecase FileUsecase
}

func New(cfg config.Config, fileUsecase FileUsecase, log *zap.Logger) *API {
	return &API{
		cfg:  cfg.Server.GRPCServer,
		addr: fmt.Sprintf(serverIPAddress, cfg.Server.GRPCServer.Port),
		log:  log,

		fileUsecase: fileUsecase,
	}
}

func (a *API) Run(ctx context.Context) error {
	return a.run(ctx)
}

func (a *API) run(ctx context.Context) error {
	a.server = grpc.NewServer(a.setOptions(ctx)...)

	svc.RegisterFileServiceServer(a.server, frontend.NewFile(a.fileUsecase, a.log))

	reflection.Register(a.server)

	a.log.Info("Service started",
		zap.String("protocol", "gRPC"),
		zap.String("address", a.addr),
	)

	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		a.log.Error("Failed to create listener", zap.Error(err))
		return fmt.Errorf("failed to create listener: %w", err)
	}

	err = a.server.Serve(listener)
	if err != nil {
		a.log.Error("gRPC server failed", zap.Error(err))
		return fmt.Errorf("failed to serve grpc: %w", err)
	}

	return nil
}
