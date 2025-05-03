package dto

import (
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/result/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ToResultFromCreateRequest(req *svc.CreateResultRequest) (model.Result, error) {
	var result model.Result
	result.UserID = req.UserId
	result.QuizID = req.QuizId
	result.Status = req.Status
	result.PassedAt = time.Now()

	for _, question := range req.Questions {
		que := model.Question{ID: question.Id}

		for _, answer := range question.Answers {
			que.Answers = append(que.Answers, model.Answer{AnswerID: answer.Id})
		}

		result.Questions = append(result.Questions, que)
	}

	return result, nil
}

func FromResultToCreateResponse(result model.Result) (*svc.CreateResultResponse, error) {

	return &svc.CreateResultResponse{
		CreatedId: result.ID,
	}, nil
}

func FromResultToGetResponse(result model.Result) (*svc.GetResultResponse, error) {
	response := &svc.GetResultResponse{
		Id:       result.ID,
		QuizId:   result.QuizID,
		UserId:   result.UserID,
		Score:    result.Score,
		Status:   result.Status,
		PassedAt: timestamppb.New(result.PassedAt),
	}

	for _, question := range result.Questions {
		que, _ := FromQuestionToGetResponse(question)
		response.Questions = append(response.Questions, que)
	}

	return response, nil
}

func FromResultToGetResponses(results []model.Result) (*svc.GetResultResponses, error) {
	var responses []*svc.GetResultResponse

	for _, result := range results {
		res, _ := FromResultToGetResponse(result)
		responses = append(responses, res)
	}

	return &svc.GetResultResponses{Results: responses}, nil
}

func FromResultToDeleteResponse(result model.Result) (*svc.DeleteResultResponse, error) {
	return &svc.DeleteResultResponse{DeletedId: result.ID}, nil
}
