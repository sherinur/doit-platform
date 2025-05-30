package frontend

import (
	"context"
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/question/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend/dto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Question struct {
	svc.UnimplementedQuestionServiceServer

	logger *zap.Logger
	uc     QuestionUseCase
}

func NewQuestion(auc QuestionUseCase, log *zap.Logger) *Question {
	return &Question{uc: auc, logger: log}
}

func (a *Question) CreateQuestion(ctx context.Context, req *svc.CreateQuestionRequest) (*svc.CreateQuestionResponse, error) {
	question, err := dto.ToQuestionFromCreateRequest(req)
	if err != nil {
		a.logger.Error("failed to parse request:", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateQuestion(ctx, question)
	if err != nil {
		a.logger.Error("failed to create question:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Created question:", zap.Any("question", resp))
	return dto.FromQuestionToCreateResponse(resp)
}

func (a *Question) CreateQuestions(ctx context.Context, req *svc.CreateQuestionRequests) (*svc.CreateQuestionResponses, error) {
	questions, err := dto.ToQuestionFromCreateRequests(req)
	if err != nil {
		a.logger.Error("failed to fetch questions request:", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateQuestions(ctx, questions)
	if err != nil {
		a.logger.Error("failed to create questions:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Created questions:", zap.Any("questions", resp))
	return dto.FromQuestionToCreateResponses(resp)
}

func (a *Question) GetQuestionById(ctx context.Context, req *svc.GetQuestionRequest) (*svc.GetQuestionResponse, error) {
	question, err := a.uc.GetQuestionById(ctx, req.Id)
	if err != nil {
		a.logger.Error("failed to get question by id:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Returned question:", zap.Any("question", question))
	return dto.FromQuestionToGetResponse(question)
}

func (a *Question) GetQuestionsByQuizId(ctx context.Context, req *svc.GetQuestionRequest) (*svc.GetQuestionResponses, error) {
	questions, err := a.uc.GetQuestionsByQuizId(ctx, req.Id)
	if err != nil {
		a.logger.Error("failed to get questions by quiz id:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Returned questions:", zap.Any("question", questions))
	return dto.FromQuestionToGetResponses(questions)
}

func (a *Question) UpdateQuestion(ctx context.Context, req *svc.UpdateQuestionRequest) (*svc.UpdateQuestionResponse, error) {
	question, err := dto.ToQuestionFromUpdateRequest(req)
	if err != nil {
		a.logger.Error("failed to fetch update question requests:", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.UpdateQuestion(ctx, question)
	if err != nil {
		a.logger.Error("failed to update question by id:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Updated question:", zap.Any("updated_id", question))
	return dto.FromQuestionToUpdateResponse(resp)
}

func (a *Question) DeleteQuestion(ctx context.Context, req *svc.DeleteQuestionRequest) (*svc.DeleteQuestionResponse, error) {
	resp, err := a.uc.DeleteQuestion(ctx, req.Id)
	if err != nil {
		a.logger.Error("failed to delete question by id:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Deleted question:", zap.Any("deleted_id", resp))
	return dto.FromQuestionToDeleteResponse(resp)
}
