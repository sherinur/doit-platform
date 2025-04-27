package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/adapters/controller/http/service/handler/dto"
)

type AnswerHandler struct {
	UseCase AnswerUseCase
}

func NewAnswerHandler(uc AnswerUseCase) *AnswerHandler {
	return &AnswerHandler{
		UseCase: uc,
	}
}

func (h *AnswerHandler) CreateAnswer(ctx *gin.Context) {
	answer, err := dto.FromAnswerCreateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	newAnswer, err := h.UseCase.CreateAnswer(ctx, answer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToAnswerResponse(newAnswer))
}

func (h *AnswerHandler) GetAnswerById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	answer, err := h.UseCase.GetAnswerById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, answer)
}

func (h *AnswerHandler) GetAnswerAll(ctx *gin.Context) {
	answers, err := h.UseCase.GetAnswerAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, answers)
}

func (h *AnswerHandler) UpdateAnswer(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	answer, err := dto.FromAnswerUpdateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	answer.ID = id
	newAnswer, err := h.UseCase.UpdateAnswer(ctx, answer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToAnswerResponse(newAnswer))
}

func (h *AnswerHandler) DeleteAnswer(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	answer, err := h.UseCase.DeleteAnswer(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToAnswerResponse(answer))
}
