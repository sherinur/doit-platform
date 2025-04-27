package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizzes-service/internal/model"
)

type QuestionRequest struct {
	Text   string `json:"text"`
	Type   string `json:"type"`
	QuizID string `json:"quiz_id"`
}

type QuestionResponse struct {
	ID string `json:"id"`
}

type QuestionGetResponse struct {
	ID      string              `json:"id"`
	Text    string              `json:"text"`
	Type    string              `json:"type"`
	QuizID  string              `json:"quiz_id"`
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
		QuizID: question.QuizID,
	}, nil
}

func ToQuestionResponse(question model.Question) QuestionResponse {
	return QuestionResponse{
		ID: question.ID,
	}
}

func ToQuestionGetResponse(question model.Question, answers []model.Answer) QuestionGetResponse {
	return QuestionGetResponse{
		ID:      question.ID,
		Text:    question.Text,
		Type:    question.Type,
		QuizID:  question.QuizID,
		Answers: ToAnswerGetAllResponse(answers),
	}
}

func ToQuestionGetAllResponse(questions []model.Question, answers []model.Answer) []QuestionGetResponse {
	questionList := make([]QuestionGetResponse, 0, len(questions))

	for _, question := range questions {
		temp := []model.Answer{}
		for _, answer := range answers {
			if answer.QuestionID == question.ID {
				temp = append(temp, answer)
			}
		}

		questionList = append(questionList, ToQuestionGetResponse(question, temp))
	}

	return questionList
}
