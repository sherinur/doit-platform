package grpcserver

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func loggingInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)

		log.Info("gRPC request completed",
			zap.String("method", info.FullMethod),
			zap.Duration("duration", time.Since(start)),
		)

		log.Debug("Request details", zap.Any("request", req))

		return resp, err
	}
}

func errorInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			log.Error("gRPC request error",
				zap.Error(err),
			)
		}

		return resp, err
	}
}

func recoveryInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("gRPC service panic recovered",
					zap.Any("panic", r),
					zap.Stack("stack"),
				)
				// err = status.Errorf(codes.Internal, "internal server error")
			}
		}()

		return handler(ctx, req)
	}
}

// TODO: Implement the tracing and metrics scraping via interceptor
// func otelInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 		resp, err := handler(ctx, req)

// 		return resp, err
// 	}
// }
