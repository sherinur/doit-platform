package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
	"net/http"
)

type QuestionRequest struct {
	Text    string          `json:"text"`
	Type    string          `json:"type"`
	Points  float64         `json:"points"`
	QuizID  string          `json:"quiz_id"`
	Answers []AnswerRequest `json:"answers"`
}

type AnswerRequest struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
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

type AnswerGetResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
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

func FromQuestionCreateRequests(ctx *gin.Context) ([]model.Question, error) {
	var requests []QuestionRequest
	err := ctx.ShouldBindJSON(&requests)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	var questions = []model.Question{}
	for _, request := range requests {
		question := model.Question{
			Text:   request.Text,
			Type:   request.Type,
			Points: request.Points,
			QuizID: request.QuizID,
		}
		questions = append(questions, question)
	}

	return questions, nil
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

func ToQuestionResponses(questions []model.Question) []QuestionResponse {
	response := []QuestionResponse{}
	for _, question := range questions {
		response = append(response, ToQuestionResponse(question))
	}

	return response
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

func ToAnswerGetResponse(answer model.Answer) AnswerGetResponse {
	return AnswerGetResponse{
		ID:   answer.AnswerID,
		Text: answer.Text,
	}
}

func ToAnswerGetAllResponse(answers []model.Answer) []AnswerGetResponse {
	answerList := make([]AnswerGetResponse, 0, len(answers))

	for _, answer := range answers {
		answerList = append(answerList, ToAnswerGetResponse(answer))
	}

	return answerList
}
