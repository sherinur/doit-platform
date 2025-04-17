package usecase

import (
	"context"
	"errors"
	"order_service/internal/adapter/http/service/handler/dto"
	"order_service/internal/model"
)

type Order struct {
	repo OrderRepo
}

func NewOrder(repo OrderRepo) *Order {
	return &Order{repo: repo}
}

func (u *Order) Create(ctx context.Context, order *model.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}

	var total float64
	for _, item := range order.OrderItems {
		total += float64(item.Quantity) * item.UnitPrice
	}
	order.TotalAmount = total
	order.Status = "pending"

	_, err := u.repo.Create(ctx, *order)

	return err
}

func (u *Order) Get(ctx context.Context, id int64) (*dto.OrderResponse, error) {
	order, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToOrderResponse(order), nil
}

func (u *Order) UpdateStatus(ctx context.Context, id int64, status string) (*dto.OrderResponse, error) {
	order, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	validStatuses := map[string]bool{
		"pending":    true,
		"processing": true,
		"completed":  true,
		"cancelled":  true,
	}
	if !validStatuses[status] {
		return nil, errors.New("invalid status")
	}

	order.Status = status
	if err := u.repo.Update(ctx, id, *order); err != nil {
		return nil, err
	}

	return dto.ToOrderResponse(order), nil
}

func (u *Order) ListByUser(ctx context.Context, userID int64) ([]dto.OrderResponse, error) {
	orders, err := u.repo.GetByUser(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.OrderResponse
	for _, order := range orders {
		responses = append(responses, *dto.ToOrderResponse(&order))
	}

	return responses, nil
}
