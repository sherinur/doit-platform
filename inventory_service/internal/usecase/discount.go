package usecase

import (
	"context"

	"inventory_service/internal/model"
)

type Discount struct {
	repo DiscountRepo
}

func NewDiscount(repo DiscountRepo) *Discount {
	return &Discount{repo: repo}
}

func (d *Discount) Create(ctx context.Context, discount model.Discount) error {
	return d.repo.Create(ctx, discount)
}

func (d *Discount) GetAllProductsWithPromotion(ctx context.Context) ([]model.Product, error) {
	return d.repo.GetAllProductsWithPromotion(ctx)
}

func (d *Discount) Delete(ctx context.Context, id int64) error {
	return d.repo.Delete(ctx, id)
}
