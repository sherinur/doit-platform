package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/model"
)

type QuestionRequest struct {
	Text      string   `json:"text"`
	Type      string   `json:"type"`
	AnswerIDs []string `json:"answer_ids"`
}

type QuestionResponse struct {
	ID string `json:"id"`
}

func FromQuestionCreateRequest(ctx *gin.Context) (model.Question, error) {
	var question QuestionRequest
	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Question{}, err
	}

	return model.Question{
		Text:      question.Text,
		Type:      question.Type,
		AnswerIDs: question.AnswerIDs,
	}, nil
}

func FromQuestionUpdateRequest(ctx *gin.Context) (model.Question, error) {
	var question QuestionRequest
	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Question{}, err
	}

	return model.Question{
		Text:      question.Text,
		Type:      question.Type,
		AnswerIDs: question.AnswerIDs,
	}, nil
}

func ToQuestionResponse(question model.Question) QuestionResponse {
	return QuestionResponse{
		ID: question.ID,
	}
}
