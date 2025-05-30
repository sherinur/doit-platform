package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/doit-platform/api-gateway/internal/model"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserID int64 `json:"user_id"`
}

func ToUserRegisterRequest(ctx *gin.Context) (model.User, error) {
	var req RegisterRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return model.User{}, err

	}

	return model.User{
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func FromUserToRegisterResponse(client model.User) RegisterResponse {
	return RegisterResponse{
		UserID: client.ID,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func ToUserLoginRequest(ctx *gin.Context) (model.User, error) {
	var req LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.User{}, err
	}
	return model.User{
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func FromTokenToLoginResponse(token model.Token) LoginResponse {
	return LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
}

type ProfileRequest struct{}

type ProfileResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func FromUserToProfileResponse(user model.User) ProfileResponse {
	return ProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Role:  user.Role,
	}
}

func ToUserUpdateProfileRequest(ctx *gin.Context) (model.UserUpdateData, error) {
	var req UpdateProfileRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.UserUpdateData{}, err
	}
	return model.UserUpdateData{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	}, nil
}

type UpdateProfileRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UpdateProfileResponse struct {
	Status string `json:"status"`
}

func FromUserToUpdateProfileResponse() UpdateProfileResponse {
	return UpdateProfileResponse{
		Status: "profile updated",
	}
}

func ToUserUpdatePasswordRequest(ctx *gin.Context) (model.UserUpdateData, error) {
	var req UpdatePasswordRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.UserUpdateData{}, err
	}
	return model.UserUpdateData{
		NewPassword: req.NewPassword,
	}, nil
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type UpdatePasswordResponse struct {
	Status string `json:"status"`
}

func FromUserToUpdatePasswordResponse() UpdatePasswordResponse {
	return UpdatePasswordResponse{
		Status: "password updated",
	}
}

func ToRefreshTokenRequest(ctx *gin.Context) (string, error) {
	var req RefreshTokenRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", err
	}
	return req.RefreshToken, nil
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func FromTokenToRefreshTokenResponse(accessToken, refreshToken string) RefreshTokenResponse {
	return RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func FromUserToDeleteAccountResponse() DeleteAccountResponse {
	return DeleteAccountResponse{
		Status: "account deleted",
	}
}

type DeleteAccountRequest struct{}

type DeleteAccountResponse struct {
	Status string `json:"status"`
}

type GetAllUsersRequest struct{}

type GetAllUsersResponse struct {
	Users []User `json:"users"`
}

func FromUsersToGetAllUsersResponse(users []*model.User) GetAllUsersResponse {
	var resUsers []User
	for _, u := range users {
		resUsers = append(resUsers, User{
			ID:    u.ID,
			Name:  u.Name,
			Phone: u.Phone,
			Email: u.Email,
			Role:  u.Role,
		})
	}
	return GetAllUsersResponse{
		Users: resUsers,
	}
}

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type ChangeUserRoleRequest struct {
	UserID  int64  `json:"user_id"`
	NewRole string `json:"new_role"`
}

type ChangeUserRoleResponse struct {
	Status string `json:"status"`
}

func ToChangeUserRoleRequest(ctx *gin.Context) (int64, string, error) {
	var req ChangeUserRoleRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, "", err
	}
	return req.UserID, req.NewRole, nil
}

func FromUserToChangeUserRoleResponse() ChangeUserRoleResponse {
	return ChangeUserRoleResponse{
		Status: "role changed",
	}
}

type VerifyEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type VerifyEmailResponse struct {
	Status string `json:"status"`
}

func ToVerifyEmailRequest(ctx *gin.Context) (string, string, error) {
	var req VerifyEmailRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", "", err
	}
	return req.Email, req.Code, nil
}

func FromUserToVerifyEmailResponse() VerifyEmailResponse {
	return VerifyEmailResponse{
		Status: "email verified",
	}
}

type SendVerificationCodeRequest struct {
	Email string `json:"email"`
}

type SendVerificationCodeResponse struct {
	Status string `json:"status"`
}

func ToSendVerificationCodeRequest(ctx *gin.Context) (string, error) {
	var req SendVerificationCodeRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", err
	}
	return req.Email, nil
}

func FromUserToSendVerificationCodeResponse() SendVerificationCodeResponse {
	return SendVerificationCodeResponse{
		Status: "code sent",
	}
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type LogoutResponse struct {
	Status string `json:"status"`
}

func ToLogoutRequest(ctx *gin.Context) (string, error) {
	var req LogoutRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", err
	}
	return req.RefreshToken, nil
}

func FromUserToLogoutResponse() LogoutResponse {
	return LogoutResponse{
		Status: "logged out",
	}
}
