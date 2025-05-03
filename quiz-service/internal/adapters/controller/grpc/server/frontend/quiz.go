package frontend

import (
	"context"
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/quiz/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Quiz struct {
	svc.UnimplementedQuizServiceServer

	uc QuizUseCase
}

func NewQuiz(auc QuizUseCase) *Quiz {
	return &Quiz{uc: auc}
}

func (a *Quiz) CreateQuiz(ctx context.Context, req *svc.CreateQuizRequest) (*svc.CreateQuizResponse, error) {
	result, err := dto.ToQuizFromCreateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateQuiz(ctx, result)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuizToCreateResponse(resp)
}

func (a *Quiz) GetQuizById(ctx context.Context, req *svc.GetQuizRequest) (*svc.GetQuizResponse, error) {
	quiz, err := a.uc.GetQuizById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuizToGetResponse(quiz)
}

func (a *Quiz) UpdateQuiz(ctx context.Context, req *svc.UpdateQuizRequest) (*svc.UpdateQuizResponse, error) {
	result, err := dto.ToQuizFromUpdateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.UpdateQuiz(ctx, result)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuizToUpdateResponse(resp)
}

func (a *Quiz) DeleteQuiz(ctx context.Context, req *svc.DeleteQuizRequest) (*svc.DeleteQuizResponse, error) {
	resp, err := a.uc.DeleteQuiz(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuizToDeleteResponse(resp)
}
