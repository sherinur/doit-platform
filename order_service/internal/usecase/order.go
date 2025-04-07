package usecase

import (
	"context"

	"order_service/internal/model"
)

type Order struct {
	repo OrderRepo
}

func NewOrder(repo OrderRepo) *Order {
	return &Order{repo: repo}
}

func (u *Order) Create(ctx context.Context, request model.Order) error {
	return u.repo.Create(ctx, request)
}

func (u *Order) Get(ctx context.Context, id uint64) (*model.Order, error) {
	return u.repo.Get(ctx, id)
}

func (u *Order) Update(ctx context.Context, id uint64, order model.Order)
