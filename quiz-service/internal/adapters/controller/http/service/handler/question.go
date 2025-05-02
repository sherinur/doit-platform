package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/http/service/handler/dto"
	"net/http"
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

	resp, err := h.UseCase.CreateQuestion(ctx, question)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToQuestionResponse(resp))
}

func (h *QuestionHandler) CreateQuestions(ctx *gin.Context) {
	questions, err := dto.FromQuestionCreateRequests(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	resp, err := h.UseCase.CreateQuestions(ctx, questions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToQuestionResponses(resp))
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

	ctx.JSON(http.StatusOK, dto.ToQuestionGetResponse(question))
}

func (h *QuestionHandler) GetQuestionsByQuizId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	questions, err := h.UseCase.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToQuestionGetAllResponse(questions))
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
