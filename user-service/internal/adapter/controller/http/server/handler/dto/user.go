package dto

import (
	"net/http"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUserFromRegisterRequest(ctx *gin.Context) (model.User, error) {
	var req RegisterRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return model.User{}, err
	}

	err = validateUserRegisterRequest(req)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Name:            req.Name,
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func FromUserToCreateResponse(user model.User) RegisterResponse {
	return RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

type LoginUserRequset struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
