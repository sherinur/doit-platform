package usecase

import (
	"context"

	"inventory_service/internal/model"
)

type ProductRepo interface {
	Create(ctx context.Context, product model.Product) error
	Get(ctx context.Context, id int64) (*model.Product, error)
	List(ctx context.Context, category string, limit, offset int) ([]model.Product, error)
	Update(ctx context.Context, id int64, updated model.Product) error
	Delete(ctx context.Context, id int64) error
}

type DiscountRepo interface {
	Create(ctx context.Context, discount model.Discount) error
	GetAllProductsWithPromotion(ctx context.Context) ([]model.Product, error)
	Delete(ctx context.Context, discountID int64) error
}
