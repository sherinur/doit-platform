package usecase

import (
	"context"
	"user-services/internal/domain/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, userID int64) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, user *model.User, userID int64) error
	Delete(ctx context.Context, userID int64) error
}
