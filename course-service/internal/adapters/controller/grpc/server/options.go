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
			MaxConnectionAge:      a.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: a.cfg.MaxConnectionAgeGrace,
		}),
		grpc.MaxRecvMsgSize(a.cfg.MaxRecvMsgSizeMiB * (1024 * 1024) /*MB*/),

		grpc.StatsHandler(otelgrpc.NewServerHandler(
			otelgrpc.WithTracerProvider(tracerProvider),
			otelgrpc.WithMeterProvider(meterProvider),
		)),

		// Interceptors
		grpc.ChainUnaryInterceptor(
			loggingInterceptor(a.log),
			errorInterceptor(a.log),
			recoveryInterceptor(a.log),
		),
	}

	return opts
}
