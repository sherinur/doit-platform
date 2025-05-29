package usecase

import (
	"context"
	"course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepo interface {
	Create(ctx context.Context, category *model.Category) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.Category, error)
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

func (uc *CategoryUsecase) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Category, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *CategoryUsecase) List(ctx context.Context) ([]*model.Category, error) {
	return uc.repo.FindAll(ctx)
}
