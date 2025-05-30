package grpcserver

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func (a *API) setOptions(ctx context.Context) []grpc.ServerOption {
	tracerProvider := otel.GetTracerProvider()
	meterProvider := otel.GetMeterProvider()

	opts := []grpc.ServerOption{
		// Params
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      a.cfg.Server.GRPCServer.MaxConnectionAge,
			MaxConnectionAgeGrace: a.cfg.Server.GRPCServer.MaxConnectionAgeGrace,
		}),
		grpc.MaxRecvMsgSize(a.cfg.Server.GRPCServer.MaxRecvMsgSizeMiB * (1024 * 1024) /*MB*/),

		grpc.StatsHandler(otelgrpc.NewServerHandler(
			otelgrpc.WithTracerProvider(tracerProvider),
			otelgrpc.WithMeterProvider(meterProvider),
		)),

		// Interceptors
		grpc.ChainUnaryInterceptor(
			AuthInterceptor(a.cfg.Jwt.JwtRefreshSecret),
			loggingInterceptor(a.log),
			errorInterceptor(a.log),
			recoveryInterceptor(a.log),
		),
	}

	return opts
}
