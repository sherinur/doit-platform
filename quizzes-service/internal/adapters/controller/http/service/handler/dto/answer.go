package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/model"
)

type AnswerRequest struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type AnswerResponse struct {
	ID string `json:"id"`
}

func FromAnswerCreateRequest(ctx *gin.Context) (model.Answer, error) {
	var answer AnswerRequest
	err := ctx.ShouldBindJSON(&answer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Answer{}, err
	}

	return model.Answer{
		Text:      answer.Text,
		IsCorrect: answer.IsCorrect,
	}, nil
}

func FromAnswerUpdateRequest(ctx *gin.Context) (model.Answer, error) {
	var answer AnswerRequest
	err := ctx.ShouldBindJSON(&answer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Answer{}, err
	}

	return model.Answer{
		Text:      answer.Text,
		IsCorrect: answer.IsCorrect,
	}, nil
}

func ToAnswerResponse(answer model.Answer) AnswerResponse {
	return AnswerResponse{
		ID: answer.ID,
	}
}
