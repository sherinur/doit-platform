package handler

import (
	"context"

	"inventory_service/internal/adapter/http/service/handler/dto"
	"inventory_service/internal/model"
)

type ProductUsecase interface {
	Create(ctx context.Context, product *model.Product) error
	Get(ctx context.Context, id int64) (*dto.ProductResponse, error)
	Update(ctx context.Context, id int64, updated *model.Product) (*dto.ProductResponse, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, category string, limit, offset int) ([]dto.ProductResponse, error)
}

type DiscountUsecase interface {
	Create(ctx context.Context, discount model.Discount) error
	GetAllProductsWithPromotion(ctx context.Context) ([]model.Product, error)
	Delete(ctx context.Context, id int64) error
}
