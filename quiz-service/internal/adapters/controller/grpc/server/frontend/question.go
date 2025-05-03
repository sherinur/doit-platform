package frontend

import (
	"context"
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/question/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Question struct {
	svc.UnimplementedQuestionServiceServer

	uc QuestionUseCase
}

func NewQuestion(auc QuestionUseCase) *Question {
	return &Question{uc: auc}
}

func (a *Question) CreateQuestion(ctx context.Context, req *svc.CreateQuestionRequest) (*svc.CreateQuestionResponse, error) {
	question, err := dto.ToQuestionFromCreateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateQuestion(ctx, question)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuestionToCreateResponse(resp)
}

func (a *Question) CreateQuestions(ctx context.Context, req *svc.CreateQuestionRequests) (*svc.CreateQuestionResponses, error) {
	questions, err := dto.ToQuestionFromCreateRequests(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateQuestions(ctx, questions)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuestionToCreateResponses(resp)
}

func (a *Question) GetQuestionById(ctx context.Context, req *svc.GetQuestionRequest) (*svc.GetQuestionResponse, error) {
	question, err := a.uc.GetQuestionById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuestionToGetResponse(question)
}

func (a *Question) GetQuestionsByQuizId(ctx context.Context, req *svc.GetQuestionRequest) (*svc.GetQuestionResponses, error) {
	questions, err := a.uc.GetQuestionsByQuizId(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuestionToGetResponses(questions)
}

func (a *Question) UpdateQuestion(ctx context.Context, req *svc.UpdateQuestionRequest) (*svc.UpdateQuestionResponse, error) {
	question, err := dto.ToQuestionFromUpdateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.UpdateQuestion(ctx, question)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuestionToUpdateResponse(resp)
}

func (a *Question) DeleteQuestion(ctx context.Context, req *svc.DeleteQuestionRequest) (*svc.DeleteQuestionResponse, error) {
	resp, err := a.uc.DeleteQuestion(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromQuestionToDeleteResponse(resp)
}
