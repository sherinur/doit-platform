package usecase

import (
	"context"
	"user-services/internal/domain/model"
)

type userUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *userUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) RegisterUser(ctx context.Context, user *model.User) error {
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) GetUserById(ctx context.Context, userID int64) (*model.User, error) {
	return u.repo.GetById(ctx, userID)
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *model.User, userID int64) error {
	return u.repo.Update(ctx, user, userID)
}

func (u *userUsecase) DeleteUser(ctx context.Context, userID int64) error {
	return u.repo.Delete(ctx, userID)
}
