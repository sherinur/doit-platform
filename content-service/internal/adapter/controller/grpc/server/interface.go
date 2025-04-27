package grpcserver

import "github.com/sherinur/doit-platform/content-service/internal/adapter/controller/grpc/server/frontend"

type FileUsecase interface {
	frontend.FileUsecase
}
