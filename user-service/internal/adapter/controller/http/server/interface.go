package httpserver

import "user-services/internal/adapter/controller/http/server/handler"

type UserUsecase interface {
	handler.UserUsecase
}
