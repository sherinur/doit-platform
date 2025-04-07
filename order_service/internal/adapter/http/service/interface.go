package service

import "order_service/internal/adapter/http/service/handler"

type OrderUsecase interface {
	handler.OrderUsecase
}
