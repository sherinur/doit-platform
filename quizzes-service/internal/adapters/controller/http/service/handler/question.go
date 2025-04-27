package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/adapters/controller/http/service/handler/dto"
)

type QuestionHandler struct {
	UseCase QuestionUseCase
}

func NewQuestionHandler(uc QuestionUseCase) *QuestionHandler {
	return &QuestionHandler{
		UseCase: uc,
	}
}

func (h *QuestionHandler) CreateQuestion(ctx *gin.Context) {
	question, err := dto.FromQuestionCreateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	newQuestion, err := h.UseCase.CreateQuestion(ctx, question)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToQuestionResponse(newQuestion))
}

func (h *QuestionHandler) GetQuestionById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	question, err := h.UseCase.GetQuestionById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) GetQuestionAll(ctx *gin.Context) {
	questions, err := h.UseCase.GetQuestionAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) UpdateQuestion(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	question, err := dto.FromQuestionUpdateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	question.ID = id
	newQuestion, err := h.UseCase.UpdateQuestion(ctx, question)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToQuestionResponse(newQuestion))
}

func (h *QuestionHandler) DeleteQuestion(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	question, err := h.UseCase.DeleteQuestion(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToQuestionResponse(question))
}
