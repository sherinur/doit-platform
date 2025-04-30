package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/model"
)

type AnswerRequest struct {
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"`
	QuestionID string `json:"question_id"`
}

type AnswerResponse struct {
	ID string `json:"id"`
}

type AnswerGetResponse struct {
	ID         string `json:"id"`
	Text       string `json:"text"`
	QuestionID string `json:"question_id"`
}

func FromAnswerCreateRequest(ctx *gin.Context) (model.Answer, error) {
	var answer AnswerRequest
	err := ctx.ShouldBindJSON(&answer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Answer{}, err
	}

	return model.Answer{
		Text:       answer.Text,
		IsCorrect:  answer.IsCorrect,
		QuestionID: answer.QuestionID,
	}, nil
}

func FromAnswerCreateRequests(ctx *gin.Context) ([]model.Answer, error) {
	var requests []AnswerRequest
	err := ctx.ShouldBindJSON(&requests)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	var answers = []model.Answer{}
	for _, request := range requests {
		answer := model.Answer{
			Text:       request.Text,
			IsCorrect:  request.IsCorrect,
			QuestionID: request.QuestionID,
		}
		answers = append(answers, answer)
	}

	return answers, nil
}

func FromAnswerUpdateRequest(ctx *gin.Context) (model.Answer, error) {
	var answer AnswerRequest
	err := ctx.ShouldBindJSON(&answer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Answer{}, err
	}

	return model.Answer{
		Text:       answer.Text,
		IsCorrect:  answer.IsCorrect,
		QuestionID: answer.QuestionID,
	}, nil
}

func ToAnswerResponse(answer model.Answer) AnswerResponse {
	return AnswerResponse{
		ID: answer.ID,
	}
}

func ToAnswerResponses(answers []model.Answer) []AnswerResponse {
	response := []AnswerResponse{}
	for _, answer := range answers {
		response = append(response, ToAnswerResponse(answer))
	}

	return response
}

func ToAnswerGetResponse(answer model.Answer) AnswerGetResponse {
	return AnswerGetResponse{
		ID:         answer.ID,
		Text:       answer.Text,
		QuestionID: answer.QuestionID,
	}
}

func ToAnswerGetAllResponse(answers []model.Answer) []AnswerGetResponse {
	answerList := make([]AnswerGetResponse, 0, len(answers))

	for _, answer := range answers {
		answerList = append(answerList, ToAnswerGetResponse(answer))
	}

	return answerList
}
