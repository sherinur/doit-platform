package dto

import (
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/answer/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

func ToAnswerFromCreateRequest(req *svc.CreateAnswerRequest) (model.Answer, error) {
	return model.Answer{
		Text:       req.Text,
		IsCorrect:  req.IsCorrect,
		QuestionID: req.QuestionId,
	}, nil
}

func FromAnswerToCreateResponse(answer model.Answer) (*svc.CreateAnswerResponse, error) {
	return &svc.CreateAnswerResponse{
		Id: answer.ID,
	}, nil
}

func ToAnswerFromCreateRequests(req *svc.CreateAnswerRequests) ([]model.Answer, error) {
	var answers []model.Answer

	for _, answer := range req.Answers {
		ans, _ := ToAnswerFromCreateRequest(answer)
		answers = append(answers, ans)
	}

	return answers, nil
}

func FromAnswersToCreateResponses(answers []model.Answer) (*svc.CreateAnswerResponses, error) {
	var responses []*svc.CreateAnswerResponse
	for _, answer := range answers {
		resp := &svc.CreateAnswerResponse{Id: answer.ID}
		responses = append(responses, resp)
	}
	return &svc.CreateAnswerResponses{
		Answers: responses,
	}, nil
}

func FromAnswerToGetResponse(answer model.Answer) (*svc.GetAnswerResponse, error) {
	return &svc.GetAnswerResponse{
		Id:         answer.ID,
		Text:       answer.Text,
		QuestionId: answer.QuestionID,
	}, nil
}

func FromAnswersToGetResponses(answers []model.Answer) (*svc.GetAnswerResponses, error) {
	var responses []*svc.GetAnswerResponse
	for _, answer := range answers {
		resp, _ := FromAnswerToGetResponse(answer)
		responses = append(responses, resp)
	}

	return &svc.GetAnswerResponses{Answers: responses}, nil
}

func ToAnswerFromUpdateRequest(req *svc.UpdateAnswerRequest) (model.Answer, error) {
	answer := req.Answer
	return model.Answer{
		ID:         answer.Id,
		Text:       answer.Text,
		QuestionID: answer.QuestionId,
		IsCorrect:  answer.IsCorrect,
	}, nil
}

func FromAnswerToUpdateResponse(answer model.Answer) (*svc.UpdateAnswerResponse, error) {
	return &svc.UpdateAnswerResponse{UpdatedId: answer.ID}, nil
}

func FromAnswerToDeleteResponse(answer model.Answer) (*svc.DeleteAnswerResponse, error) {
	return &svc.DeleteAnswerResponse{DeletedId: answer.ID}, nil
}
