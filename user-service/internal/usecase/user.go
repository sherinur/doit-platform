package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
	"github.com/sherinur/doit-platform/user-service/pkg/security"
)

type userUsecase struct {
	userRepo        UserRepo
	tokenRepo       RefreshTokenRepo
	jwtManager      *security.JWTManager
	passwordManager *security.PasswordManager
}

func NewUserUsecase(
	userRepo UserRepo,
	tokenRepo RefreshTokenRepo,
	jwtManager *security.JWTManager,
	passwordManager *security.PasswordManager,
) *userUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		tokenRepo:       tokenRepo,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
	}
}

func (uc *userUsecase) RegisterUser(ctx context.Context, request *model.User) (*model.User, error) {
	// Validate the user input
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Check if the user already exists
	existingUser, err := uc.userRepo.GetByEmail(ctx, request.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if existingUser != nil {
		return nil, model.ErrUserExists
	}

	// Hash the password
	request.PasswordHash, err = uc.passwordManager.HashPassword(request.CurrentPassword)
	if err != nil {
		return nil, err
	}

	request.CreatedAt = time.Now().UTC()
	request.UpdatedAt = time.Now().UTC()
	request.Role = "user"

	newUser, err := uc.userRepo.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (uc *userUsecase) LoginUser(ctx context.Context, request *model.User) (model.Token, error) {
	err := request.Validate()
	if err != nil {
		return model.Token{}, err
	}

	user, err := uc.userRepo.GetByEmail(ctx, request.Email)
	if err != nil {
		return model.Token{}, err
	}

	err = uc.passwordManager.CheckPassword(user.PasswordHash, request.CurrentPassword)
	if err != nil {
		return model.Token{}, err
	}

	accessPayload := uc.jwtManager.CreateAccessPayload(user)
	refreshPayload := uc.jwtManager.CreateRefreshPayload(user)

	accessToken, refreshToken, err := uc.jwtManager.GenerateTokens(accessPayload, refreshPayload)
	if err != nil {
		return model.Token{}, err
	}

	session := model.Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	err = uc.tokenRepo.Create(ctx, &session)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *userUsecase) RefreshToken(ctx context.Context, refreshToken string) (model.Token, error) {
	session, err := uc.tokenRepo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return model.Token{}, err
	}
	if session.ExpiresAt.Before(time.Now().UTC()) {
		return model.Token{}, model.ErrRefreshTokenExpired
	}

	user, err := uc.userRepo.GetById(ctx, session.UserID)
	if err != nil {
		return model.Token{}, err
	}

	accessPayload := uc.jwtManager.CreateAccessPayload(user)
	refreshPayload := uc.jwtManager.CreateRefreshPayload(user)

	accessToken, newRefreshToken, err := uc.jwtManager.GenerateTokens(accessPayload, refreshPayload)
	if err != nil {
		return model.Token{}, err
	}

	// delete old refresh and insert new one (rotation)
	err = uc.tokenRepo.DeleteByRefreshToken(ctx, refreshToken)
	if err != nil {
		return model.Token{}, err
	}

	newSession := model.Session{
		UserID:       user.ID,
		RefreshToken: newRefreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	err = uc.tokenRepo.Create(ctx, &newSession)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *userUsecase) GetUserById(ctx context.Context, userID int64) (*model.User, error) {
	return uc.userRepo.GetById(ctx, userID)
}

func (uc *userUsecase) UpdateUser(ctx context.Context, user *model.User, userID int64) error {
	return uc.userRepo.Update(ctx, user, userID)
}

func (uc *userUsecase) DeleteUser(ctx context.Context, userID int64) error {
	return uc.userRepo.Delete(ctx, userID)
}
