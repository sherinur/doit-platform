package usecase

import (
	"context"
	"course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagRepo interface {
	Create(ctx context.Context, tag *model.Tag) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.Tag, error)
	FindAll(ctx context.Context) ([]*model.Tag, error)
}

type TagUsecase struct {
	repo TagRepo
}

func NewTagUsecase(repo TagRepo) *TagUsecase {
	return &TagUsecase{repo: repo}
}

func (uc *TagUsecase) Create(ctx context.Context, tag *model.Tag) error {
	return uc.repo.Create(ctx, tag)
}

func (uc *TagUsecase) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Tag, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *TagUsecase) List(ctx context.Context) ([]*model.Tag, error) {
	return uc.repo.FindAll(ctx)
}
