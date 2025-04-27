package grpcserver

import "content-service/internal/adapter/controller/grpc/server/frontend"

type FileUsecase interface {
	frontend.FileUsecase
}
