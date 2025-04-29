package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/model"
)

type QuestionRequest struct {
	Text   string  `json:"text"`
	Type   string  `json:"type"`
	Points float64 `json:"points"`
	QuizID string  `json:"quiz_id"`
}

type QuestionResponse struct {
	ID string `json:"id"`
}

type QuestionGetResponse struct {
	ID      string              `json:"id"`
	Text    string              `json:"text"`
	Type    string              `json:"type"`
	QuizID  string              `json:"quiz_id"`
	Points  float64             `json:"points"`
	Answers []AnswerGetResponse `json:"answers"`
}

func FromQuestionCreateRequest(ctx *gin.Context) (model.Question, error) {
	var question QuestionRequest
	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Question{}, err
	}

	return model.Question{
		Text:   question.Text,
		Type:   question.Type,
		Points: question.Points,
		QuizID: question.QuizID,
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
		Text:   question.Text,
		Type:   question.Type,
		Points: question.Points,
		QuizID: question.QuizID,
	}, nil
}

func ToQuestionResponse(question model.Question) QuestionResponse {
	return QuestionResponse{
		ID: question.ID,
	}
}

func ToQuestionGetResponse(question model.Question) QuestionGetResponse {
	return QuestionGetResponse{
		ID:      question.ID,
		Text:    question.Text,
		Type:    question.Type,
		Points:  question.Points,
		QuizID:  question.QuizID,
		Answers: ToAnswerGetAllResponse(question.Answers),
	}
}

func ToQuestionGetAllResponse(questions []model.Question) []QuestionGetResponse {
	questionList := make([]QuestionGetResponse, 0, len(questions))

	for _, question := range questions {
		questionList = append(questionList, ToQuestionGetResponse(question))
	}

	return questionList
}
