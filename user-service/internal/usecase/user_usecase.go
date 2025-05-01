package usecase

import (
	"context"
	"time"

	"user-services/internal/domain/model"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *userUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) RegisterUser(ctx context.Context, request *model.User) (*model.User, error) {
	// check for uniqueness
	existingUser, err := uc.repo.GetByEmail(ctx, request.Email)
	if err != nil {
		return &model.User{}, err
	} else if existingUser != nil {
		return &model.User{}, model.ErrUserExists
	}

	request.PasswordHash, err = uc.hashPassword(request.CurrentPassword)
	if err != nil {
		return &model.User{}, err
	}

	request.CreatedAt = time.Now().UTC()
	request.UpdatedAt = time.Now().UTC()

	newuser, err := uc.repo.Create(ctx, request)
	if err != nil {
		return &model.User{}, err
	}
	request.ID = newuser.ID

	return request, nil
}

func (uc *userUsecase) GetUserById(ctx context.Context, userID int64) (*model.User, error) {
	return uc.repo.GetById(ctx, userID)
}

func (uc *userUsecase) UpdateUser(ctx context.Context, user *model.User, userID int64) error {
	return uc.repo.Update(ctx, user, userID)
}

func (uc *userUsecase) DeleteUser(ctx context.Context, userID int64) error {
	return uc.repo.Delete(ctx, userID)
}

func (uc *userUsecase) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
