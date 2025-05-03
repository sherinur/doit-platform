package dto

import (
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/question/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

func ToQuestionFromCreateRequest(req *svc.CreateQuestionRequest) (model.Question, error) {
	var question model.Question
	question.Text = req.Text
	question.Type = req.Type
	question.Points = req.Points
	question.QuizID = req.QuizId

	for _, answer := range req.Answers {
		question.Answers = append(question.Answers, model.Answer{AnswerID: answer.AnswerId, Text: answer.Text, IsCorrect: answer.IsCorrect})
	}
	return question, nil
}

func FromQuestionToCreateResponse(question model.Question) (*svc.CreateQuestionResponse, error) {
	return &svc.CreateQuestionResponse{
		CreatedId: question.ID,
	}, nil
}

func ToQuestionFromCreateRequests(req *svc.CreateQuestionRequests) ([]model.Question, error) {
	var questions []model.Question

	for _, question := range req.Questions {
		que, _ := ToQuestionFromCreateRequest(question)
		questions = append(questions, que)
	}

	return questions, nil
}

func FromQuestionToCreateResponses(questions []model.Question) (*svc.CreateQuestionResponses, error) {
	var responses []*svc.CreateQuestionResponse
	for _, question := range questions {
		resp := &svc.CreateQuestionResponse{CreatedId: question.ID}
		responses = append(responses, resp)
	}
	return &svc.CreateQuestionResponses{
		Questions: responses,
	}, nil
}

func FromQuestionToGetResponse(question model.Question) (*svc.GetQuestionResponse, error) {
	var answers []*svc.GetAnswerResponse

	for _, answer := range question.Answers {
		answers = append(answers, &svc.GetAnswerResponse{Id: answer.AnswerID, Text: answer.Text})
	}

	return &svc.GetQuestionResponse{
		Id:      question.ID,
		Text:    question.Text,
		Type:    question.Type,
		Points:  question.Points,
		QuizId:  question.QuizID,
		Answers: answers,
	}, nil
}

func FromQuestionToGetResponses(questions []model.Question) (*svc.GetQuestionResponses, error) {
	var responses []*svc.GetQuestionResponse
	for _, question := range questions {
		resp, _ := FromQuestionToGetResponse(question)
		responses = append(responses, resp)
	}

	return &svc.GetQuestionResponses{Questions: responses}, nil
}

func ToQuestionFromUpdateRequest(req *svc.UpdateQuestionRequest) (model.Question, error) {
	question := req.Question
	var response model.Question

	response.ID = question.Id
	response.Text = question.Text
	response.Points = question.Points
	response.Type = question.Type
	response.QuizID = question.QuizId

	for _, answer := range question.Answers {
		response.Answers = append(response.Answers, model.Answer{AnswerID: answer.AnswerId, Text: answer.Text, IsCorrect: answer.IsCorrect})
	}

	return response, nil
}

func FromQuestionToUpdateResponse(question model.Question) (*svc.UpdateQuestionResponse, error) {
	return &svc.UpdateQuestionResponse{UpdatedId: question.ID}, nil
}

func FromQuestionToDeleteResponse(question model.Question) (*svc.DeleteQuestionResponse, error) {
	return &svc.DeleteQuestionResponse{DeletedId: question.ID}, nil
}
