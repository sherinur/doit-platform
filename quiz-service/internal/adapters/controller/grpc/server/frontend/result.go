package frontend

import (
	"context"
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/result/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend/dto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Result struct {
	svc.UnimplementedResultServiceServer

	logger *zap.Logger
	uc     ResultUseCase
}

func NewResult(auc ResultUseCase, log *zap.Logger) *Result {
	return &Result{uc: auc, logger: log}
}

func (a *Result) CreateResult(ctx context.Context, req *svc.CreateResultRequest) (*svc.CreateResultResponse, error) {
	result, err := dto.ToResultFromCreateRequest(req)
	if err != nil {
		a.logger.Error("failed to parse create request:", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateResult(ctx, result)
	if err != nil {
		a.logger.Error("failed to create request:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Created result:", zap.Any("result", resp))
	return dto.FromResultToCreateResponse(resp)
}

func (a *Result) GetResultById(ctx context.Context, req *svc.GetResultRequest) (*svc.GetResultResponse, error) {
	result, err := a.uc.GetResultById(ctx, req.Id)
	if err != nil {
		a.logger.Error("failed to get result:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Returned result:", zap.Any("result", result))
	return dto.FromResultToGetResponse(result)
}

func (a *Result) GetResultsByQuizId(ctx context.Context, req *svc.GetResultRequest) (*svc.GetResultResponses, error) {
	results, err := a.uc.GetResultsByQuizId(ctx, req.Id)
	if err != nil {
		a.logger.Error("failed to get results:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Returned results:", zap.Any("result", results))
	return dto.FromResultToGetResponses(results)
}

func (a *Result) GetResultsByUserId(ctx context.Context, req *svc.GetResultRequest) (*svc.GetResultResponses, error) {
	results, err := a.uc.GetResultsByUserId(ctx, req.Id)
	if err != nil {
		a.logger.Error("failed to get results:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Returned results:", zap.Any("result", results))
	return dto.FromResultToGetResponses(results)
}

func (a *Result) DeleteResult(ctx context.Context, req *svc.DeleteResultRequest) (*svc.DeleteResultResponse, error) {
	resp, err := a.uc.DeleteResult(ctx, req.Id)
	if err != nil {
		a.logger.Error("failed to delete result:", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	a.logger.Info("Deleted result:", zap.Any("deleted_id", resp))
	return dto.FromResultToDeleteResponse(resp)
}
