package handler

import (
	"context"

	"order_service/internal/model"
)

type OrderUsecase interface {
	Create(ctx context.Context, request model.Order) error
}
