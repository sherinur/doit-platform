package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/adapters/controller/http/service/handler/dto"
)

type ResultHandler struct {
	UseCase ResultUseCase
}

func NewResultHandler(uc ResultUseCase) *ResultHandler {
	return &ResultHandler{
		UseCase: uc,
	}
}

func (h *ResultHandler) CreateResult(ctx *gin.Context) {
	result, err := dto.FromResultCreateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	newResult, err := h.UseCase.CreateResult(ctx, result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToResultResponse(newResult))
}

func (h *ResultHandler) GetResultById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	result, err := h.UseCase.GetResultById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToResultGetResponse(result))
}

func (h *ResultHandler) GetResultsByQuizId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	results, err := h.UseCase.GetResultsByQuizId(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToResultGetAllResponse(results))
}

func (h *ResultHandler) GetResultsByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	results, err := h.UseCase.GetResultsByUserId(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToResultGetAllResponse(results))
}

func (h *ResultHandler) DeleteResult(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	quiz, err := h.UseCase.DeleteResult(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToResultResponse(quiz))
}
