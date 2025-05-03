package frontend

import (
	"context"
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/result/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Result struct {
	svc.UnimplementedResultServiceServer

	uc ResultUseCase
}

func NewResult(auc ResultUseCase) *Result {
	return &Result{uc: auc}
}

func (a *Result) CreateResult(ctx context.Context, req *svc.CreateResultRequest) (*svc.CreateResultResponse, error) {
	result, err := dto.ToResultFromCreateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := a.uc.CreateResult(ctx, result)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromResultToCreateResponse(resp)
}

func (a *Result) GetResultById(ctx context.Context, req *svc.GetResultRequest) (*svc.GetResultResponse, error) {
	results, err := a.uc.GetResultById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromResultToGetResponse(results)
}

func (a *Result) GetResultsByQuizId(ctx context.Context, req *svc.GetResultRequest) (*svc.GetResultResponses, error) {
	results, err := a.uc.GetResultsByQuizId(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromResultToGetResponses(results)
}

func (a *Result) GetResultsByUserId(ctx context.Context, req *svc.GetResultRequest) (*svc.GetResultResponses, error) {
	results, err := a.uc.GetResultsByUserId(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromResultToGetResponses(results)
}

func (a *Result) DeleteResult(ctx context.Context, req *svc.DeleteResultRequest) (*svc.DeleteResultResponse, error) {
	resp, err := a.uc.DeleteResult(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromResultToDeleteResponse(resp)
}
