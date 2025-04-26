package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/model"
	"time"
)

type QuizRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CreatedBy   string   `json:"created_by"`
	Status      string   `json:"status"`
	QuestionIDs []string `json:"question_ids"`
}

type QuizResponse struct {
	ID string `json:"id"`
}

func FromQuizCreateRequest(ctx *gin.Context) (model.Quiz, error) {
	var quiz QuizRequest
	err := ctx.ShouldBindJSON(&quiz)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Quiz{}, err
	}

	return model.Quiz{
		Title:       quiz.Title,
		Description: quiz.Description,
		CreatedBy:   quiz.CreatedBy,
		Status:      quiz.Status,
		QuestionIDs: quiz.QuestionIDs,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func FromQuizUpdateRequest(ctx *gin.Context) (model.Quiz, error) {
	var quiz QuizRequest
	err := ctx.ShouldBindJSON(&quiz)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Quiz{}, err
	}

	return model.Quiz{
		Title:       quiz.Title,
		Description: quiz.Description,
		CreatedBy:   quiz.CreatedBy,
		Status:      quiz.Status,
		QuestionIDs: quiz.QuestionIDs,
		UpdatedAt:   time.Now(),
	}, nil
}

func ToQuizResponse(quiz model.Quiz) QuizResponse {
	return QuizResponse{
		ID: quiz.ID,
	}
}
