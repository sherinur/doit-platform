package server

import "github.com/sherinur/doit-platform/user-service/internal/adapter/controller/grpc/server/frontend"

type UserUsecase interface {
	frontend.UserUsecase
}
