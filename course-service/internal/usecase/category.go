package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/course-service/internal/model"
)

type CategoryRepo interface {
	Create(ctx context.Context, category *model.Category) error
	FindByID(ctx context.Context, id string) (*model.Category, error)
	FindAll(ctx context.Context) ([]*model.Category, error)
}

type CategoryUsecase struct {
	repo CategoryRepo
}

func NewCategoryUsecase(repo CategoryRepo) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (uc *CategoryUsecase) Create(ctx context.Context, category *model.Category) error {
	return uc.repo.Create(ctx, category)
}

func (uc *CategoryUsecase) GetByID(ctx context.Context, id string) (*model.Category, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *CategoryUsecase) List(ctx context.Context) ([]*model.Category, error) {
	return uc.repo.FindAll(ctx)
}
