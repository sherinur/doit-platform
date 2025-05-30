package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/doit-platform/api-gateway/internal/adapter/http/server/handler/dto"
)

type User struct {
	uc UserUsecase
}

func NewUser(uc UserUsecase) *User {
	return &User{
		uc: uc,
	}
}

func (c *User) Register(ctx *gin.Context) {
	User, err := dto.ToUserRegisterRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	newUser, err := c.uc.RegisterUser(ctx.Request.Context(), &User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromUserToRegisterResponse(*newUser))
}

func (c *User) Login(ctx *gin.Context) {
	user, err := dto.ToUserLoginRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	token, err := c.uc.LoginUser(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromTokenToLoginResponse(token))
}

func (c *User) RefreshToken(ctx *gin.Context) {
	refreshToken, err := dto.ToRefreshTokenRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	token, err := c.uc.RefreshToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromTokenToLoginResponse(token))
}

func (c *User) Logout(ctx *gin.Context) {
	refreshToken, err := dto.ToLogoutRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	err = c.uc.Logout(ctx.Request.Context(), refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *User) Profile(ctx *gin.Context) {
	user, err := c.uc.GetUserById(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromUserToProfileResponse(*user))
}

func (c *User) UpdateUserInfo(ctx *gin.Context) {
	userUpdateData, err := dto.ToUserUpdateProfileRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	err = c.uc.UpdateUserInfo(ctx.Request.Context(), &userUpdateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromUserToUpdateProfileResponse())
}

func (c *User) UpdateUserPassword(ctx *gin.Context) {
	userUpdateData, err := dto.ToUserUpdatePasswordRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}
	err = c.uc.UpdateUserPassword(ctx.Request.Context(), &userUpdateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dto.FromUserToUpdateProfileResponse())
}

func (c *User) DeleteAccount(ctx *gin.Context) {
	err := c.uc.DeleteUser(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *User) GetAllUsers(ctx *gin.Context) {
	users, err := c.uc.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.FromUsersToGetAllUsersResponse(users))
}

func (c *User) ChangeUserRole(ctx *gin.Context) {
	userID, newRole, err := dto.ToChangeUserRoleRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	err = c.uc.ChangeUserRole(ctx.Request.Context(), userID, newRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User role updated successfully"})
}

func (c *User) SendVerificationCode(ctx *gin.Context) {
	email, err := dto.ToSendVerificationCodeRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	err = c.uc.SendVerificationCode(ctx.Request.Context(), email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *User) VerifyEmail(ctx *gin.Context) {
	email, code, err := dto.ToVerifyEmailRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	err = c.uc.VerifyEmail(ctx.Request.Context(), email, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.FromUserToVerifyEmailResponse())
}
