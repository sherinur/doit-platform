package server

import (
	"context"

	"github.com/sherinur/doit-platform/user-service/internal/adapter/controller/grpc/server/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func (a *API) setOptions(ctx context.Context, secretket string) []grpc.ServerOption {
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      a.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: a.cfg.MaxConnectionAgeGrace,
		}),
		grpc.UnaryInterceptor(interceptor.AuthInterceptor(secretket)),
		grpc.MaxRecvMsgSize(a.cfg.MaxRecvMsgSizeMiB * (1024 * 1024)), // MaxRecvSize * 1 MB
	}

	return opts
}
