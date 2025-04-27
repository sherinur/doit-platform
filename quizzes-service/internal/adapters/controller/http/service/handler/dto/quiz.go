package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/model"
	"time"
)

type QuizRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	Status      string `json:"status"`
}

type QuizResponse struct {
	ID string `json:"id"`
}

type QuizGetResponse struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	CreatedBy   string                `json:"created_by"`
	Status      string                `json:"status"`
	Questions   []QuestionGetResponse `json:"questions"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
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
		UpdatedAt:   time.Now(),
	}, nil
}

func ToQuizResponse(quiz model.Quiz) QuizResponse {
	return QuizResponse{
		ID: quiz.ID,
	}
}

func ToQuizGetResponse(quiz model.Quiz, question []model.Question, answers []model.Answer) QuizGetResponse {
	return QuizGetResponse{
		ID:          quiz.ID,
		Title:       quiz.Title,
		Description: quiz.Description,
		CreatedBy:   quiz.CreatedBy,
		Status:      quiz.Status,
		Questions:   ToQuestionGetAllResponse(question, answers),
		CreatedAt:   quiz.CreatedAt,
		UpdatedAt:   quiz.UpdatedAt,
	}
}
