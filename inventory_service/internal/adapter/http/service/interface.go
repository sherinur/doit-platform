package service

import "inventory_service/internal/adapter/http/service/handler"

type ProductUsecase interface {
	handler.ProductUsecase
}

type DiscountUsecase interface {
	handler.DiscountUsecase
}
