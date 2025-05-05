package handler

import (
	"net/http"

	"github.com/sherinur/doit-platform/user-service/internal/adapter/controller/http/server/handler/dto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc UserUsecase
}

func NewUserHandler(uc UserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	user, err := dto.ToUserFromRegisterRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	newuser, err := h.uc.RegisterUser(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromUserToCreateResponse(*newuser))
}
