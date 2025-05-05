// internal/adapter/controller/grpc/server/interceptor/auth.go
package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey string

const ContextUserID ctxKey = "userID"

func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md["authorization"]
		if len(authHeader) == 0 || !strings.HasPrefix(authHeader[0], "Bearer ") {
			return nil, status.Error(codes.Unauthenticated, "invalid or missing auth token")
		}

		// tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
		// claims, err := jwt.ValidateAccessToken(tokenStr)
		// if err != nil {
		// 	return nil, status.Error(codes.Unauthenticated, "invalid token")
		// }

		// ctx = context.WithValue(ctx, ContextUserID, claims.UserID)
		return handler(ctx, req)
	}
}
