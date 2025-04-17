package usecase

import (
	"context"

	"inventory_service/internal/adapter/http/service/handler/dto"
	"inventory_service/internal/model"
)

type Product struct {
	repo ProductRepo
}

func NewProduct(repo ProductRepo) *Product {
	return &Product{repo: repo}
}

func (u *Product) Create(ctx context.Context, product *model.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	return u.repo.Create(ctx, *product)
}

func (u *Product) Get(ctx context.Context, id int64) (*dto.ProductResponse, error) {
	product, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToProductResponse(product), nil
}

func (u *Product) Update(ctx context.Context, id int64, updated *model.Product) (*dto.ProductResponse, error) {
	existing, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	existing.Name = updated.Name
	existing.Description = updated.Description
	existing.Price = updated.Price
	existing.Category = updated.Category

	if err := existing.Validate(); err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, id, *existing); err != nil {
		return nil, err
	}

	return dto.ToProductResponse(existing), nil
}

func (u *Product) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *Product) List(ctx context.Context, category string, limit, offset int) ([]dto.ProductResponse, error) {
	products, err := u.repo.List(ctx, category, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []dto.ProductResponse
	for _, product := range products {
		responses = append(responses, *dto.ToProductResponse(&product))
	}

	return responses, nil
}
