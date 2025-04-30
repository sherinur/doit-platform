package grpcserver

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func (a *API) setOptions(ctx context.Context) []grpc.ServerOption {
	opts := []grpc.ServerOption{
		// Params
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      a.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: a.cfg.MaxConnectionAgeGrace,
		}),
		grpc.MaxRecvMsgSize(a.cfg.MaxRecvMsgSizeMiB * (1024 * 1024) /*MB*/),

		// Interceptors
		grpc.ChainUnaryInterceptor(
			loggingInterceptor(a.log),
			errorInterceptor(a.log),
			recoveryInterceptor(a.log),
		),
	}

	return opts
}
