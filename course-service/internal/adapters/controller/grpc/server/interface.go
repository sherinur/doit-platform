package grpcserver

import "github.com/sherinur/doit-platform/course-service/internal/adapters/controller/grpc/server/frontend"

type FileUsecase interface {
	frontend.FileUsecase
}
