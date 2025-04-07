package usecase

import (
	"context"

	"order_service/internal/model"
)

type OrderRepo interface {
	Create(ctx context.Context, item model.Order) error
	Get(ctx context.Context, id uint64) (*model.Order, error)
	GetAll(ctx context.Context) ([]model.Order, error)
	Update(ctx context.Context, id uint64, order model.Order) error
	Delete(ctx context.Context, id uint64) error
}
