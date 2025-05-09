package dto

import (
	svc "github.com/sherinur/doit-platform/apis/gen/user-service/service/frontend/user/v1"
	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

func ToUserFromRegisterRequest(req *svc.RegisterRequest) (model.User, error) {
	return model.User{
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func FromUserToRegisterResponse(user model.User) (*svc.RegisterResponse, error) {
	return &svc.RegisterResponse{
		UserId: user.ID,
	}, nil
}

func ToUserFromLoginRequest(req *svc.LoginRequest) (model.User, error) {
	return model.User{
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func FromTokenToLoginResponse(token model.Token) (*svc.LoginResponse, error) {
	return &svc.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
