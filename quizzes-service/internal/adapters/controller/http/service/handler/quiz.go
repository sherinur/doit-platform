package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/adapters/controller/http/service/handler/dto"
)

type QuizHandler struct {
	UseCase QuizUseCase
}

func NewQuizHandler(uc QuizUseCase) *QuizHandler {
	return &QuizHandler{
		UseCase: uc,
	}
}

func (h *QuizHandler) CreateQuiz(ctx *gin.Context) {
	quiz, err := dto.FromQuizCreateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	newQuiz, err := h.UseCase.CreateQuiz(ctx, quiz)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToQuizResponse(newQuiz))
}

func (h *QuizHandler) GetQuizById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	quiz, err := h.UseCase.GetQuizById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToQuizGetResponse(quiz))
}

func (h *QuizHandler) UpdateQuiz(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	quiz, err := dto.FromQuizUpdateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	quiz.ID = id
	newQuiz, err := h.UseCase.UpdateQuiz(ctx, quiz)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToQuizResponse(newQuiz))
}

func (h *QuizHandler) DeleteQuiz(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	quiz, err := h.UseCase.DeleteQuiz(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToQuizResponse(quiz))
}
