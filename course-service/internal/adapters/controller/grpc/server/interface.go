package grpcserver

import "github.com/sherinur/doit-platform/course-service/internal/adapters/controller/grpc/server/frontend"

type CourseUsecase interface {
	frontend.CourseUsecase
}
