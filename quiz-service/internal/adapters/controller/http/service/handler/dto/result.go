package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
	"net/http"
	"time"
)

type ResultRequest struct {
	UserID    string           `json:"user_id"`
	QuizID    string           `json:"quiz_id"`
	Status    string           `json:"status"`
	Questions []ResultQuestion `json:"questions"`
}

type ResultQuestion struct {
	ID      string         `json:"id"`
	Answers []ResultAnswer `json:"answers"`
}

type ResultAnswer struct {
	ID string `json:"id"`
}

type ResultResponse struct {
	ID string `json:"id"`
}

type ResultGetResponse struct {
	ID        string                `json:"id"`
	UserID    string                `json:"user_id"`
	QuizID    string                `json:"quiz_id"`
	Score     float64               `json:"score"`
	Questions []QuestionGetResponse `json:"questions"`
	PassedAt  time.Time             `json:"passed_at"`
}

func FromResultCreateRequest(ctx *gin.Context) (model.Result, error) {
	var request ResultRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Result{}, err
	}

	var result model.Result
	result.Status = request.Status
	result.UserID = request.UserID
	result.QuizID = request.QuizID
	result.PassedAt = time.Now()

	for _, q := range request.Questions {
		question := model.Question{}
		question.ID = q.ID
		for _, answer := range q.Answers {
			question.Answers = append(question.Answers, model.Answer{ID: answer.ID})
		}
		result.Questions = append(result.Questions, question)
	}

	return result, nil
}

func ToResultResponse(result model.Result) ResultResponse {
	return ResultResponse{
		ID: result.ID,
	}
}

func ToResultGetResponse(result model.Result) ResultGetResponse {
	return ResultGetResponse{
		ID:        result.ID,
		UserID:    result.UserID,
		QuizID:    result.QuizID,
		Score:     result.Score,
		Questions: ToQuestionGetAllResponse(result.Questions),
		PassedAt:  result.PassedAt,
	}
}

func ToResultGetAllResponse(results []model.Result) []ResultGetResponse {
	resultList := make([]ResultGetResponse, 0, len(results))

	for _, result := range results {
		resultList = append(resultList, ToResultGetResponse(result))
	}

	return resultList
}
