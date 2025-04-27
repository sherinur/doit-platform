package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/sherinur/doit-platform/content-service/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverIPAddress = "0.0.0.0:%d"

type API struct {
	server *grpc.Server
	cfg    config.GRPCServer
	addr   string

	fileUsecase FileUsecase
}

func New(cfg config.Config, fileUsecase FileUsecase) *API {
	return &API{
		cfg:  cfg.Server.GRPCServer,
		addr: fmt.Sprintf(serverIPAddress, cfg.Server.GRPCServer.Port),

		fileUsecase: fileUsecase,
	}
}

func (a *API) Run(ctx context.Context) error {
	return a.run(ctx)
}

func (a *API) run(ctx context.Context) error {
	a.server = grpc.NewServer(a.setOptions(ctx)...)

	reflection.Register(a.server)

	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}

	err = a.server.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to serve grpc: %w", err)
	}

	return nil
}
