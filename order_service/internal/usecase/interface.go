package usecase

import (
	"context"

	"order_service/internal/model"
)

type OrderRepo interface {
	Create(ctx context.Context, order model.Order) (int64, error)
	Get(ctx context.Context, id int64) (*model.Order, error)
	GetAll(ctx context.Context) ([]model.Order, error)
	GetByUser(ctx context.Context) ([]model.Order, error)
	Update(ctx context.Context, id int64, updated model.Order) error
	Delete(ctx context.Context, id int64) error
}
