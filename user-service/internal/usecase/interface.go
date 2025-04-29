package usecase

import (
	"context"
	"user-services/internal/domain/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, user_id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User, user_id int64) error
	Delete(ctx context.Context, user_id int64) error
}
