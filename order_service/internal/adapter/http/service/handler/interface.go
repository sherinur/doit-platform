package handler

import (
	"context"

	"order_service/internal/adapter/http/service/handler/dto"
	"order_service/internal/model"
)

type OrderUsecase interface {
	Create(ctx context.Context, order *model.Order) error
	Get(ctx context.Context, id int64) (*dto.OrderResponse, error)
	UpdateStatus(ctx context.Context, id int64, status string) (*dto.OrderResponse, error)
	ListByUser(ctx context.Context, userID int64) ([]dto.OrderResponse, error)
}
