package frontend

import (
	"context"
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/answer/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Answer struct {
	svc.UnimplementedAnswerServiceServer

	uc AnswerUseCase
}

func NewAnswer(auc AnswerUseCase) *Answer {
	return &Answer{uc: auc}
}

func (a *Answer) CreateAnswer(ctx context.Context, req *svc.CreateAnswerRequest) (*svc.CreateAnswerResponse, error) {
	answer, err := dto.ToAnswerFromCreateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateAnswer(ctx, answer)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromAnswerToCreateResponse(resp)
}

func (a *Answer) CreateAnswers(ctx context.Context, req *svc.CreateAnswerRequests) (*svc.CreateAnswerResponses, error) {
	answers, err := dto.ToAnswerFromCreateRequests(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateAnswers(ctx, answers)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromAnswersToCreateResponses(resp)
}

func (a *Answer) GetAnswerById(ctx context.Context, req *svc.GetAnswerRequest) (*svc.GetAnswerResponse, error) {
	answer, err := a.uc.GetAnswerById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromAnswerToGetResponse(answer)
}

func (a *Answer) GetAnswersByQuestionId(ctx context.Context, req *svc.GetAnswerRequest) (*svc.GetAnswerResponses, error) {
	answers, err := a.uc.GetAnswersByQuestionId(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromAnswersToGetResponses(answers)
}

func (a *Answer) UpdateAnswer(ctx context.Context, req *svc.UpdateAnswerRequest) (*svc.UpdateAnswerResponse, error) {
	answer, err := dto.ToAnswerFromUpdateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.UpdateAnswer(ctx, answer)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromAnswerToUpdateResponse(resp)
}

func (a *Answer) DeleteAnswer(ctx context.Context, req *svc.DeleteAnswerRequest) (*svc.DeleteAnswerResponse, error) {
	resp, err := a.uc.DeleteAnswer(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromAnswerToDeleteResponse(resp)
}
