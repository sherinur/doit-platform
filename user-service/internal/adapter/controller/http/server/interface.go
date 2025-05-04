package httpserver

import "github.com/sherinur/doit-platform/user-service/internal/adapter/controller/http/server/handler"

type UserUsecase interface {
	handler.UserUsecase
}
