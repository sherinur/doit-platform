package frontend

import (
	"context"

	feedbacksvc "github.com/sherinur/doit-platform/apis/gen/social-service/service/frontend/feedback/v1"
	"github.com/sherinur/doit-platform/social-service/internal/adapter/grpc/server/frontend/dto"
	"go.uber.org/zap"
)

type Feedback struct {
	feedbacksvc.UnimplementedFeedbackServiceServer

	log *zap.Logger
	uc  FeedbackUsecase
}

func NewFeedback(uc FeedbackUsecase, logger *zap.Logger) *Feedback {
	return &Feedback{
		log: logger,
		uc:  uc,
	}
}

func (f *Feedback) AddFeedback(ctx context.Context, req *feedbacksvc.AddFeedbackRequest) (*feedbacksvc.AddFeedbackResponse, error) {
	fb := dto.ToModelFromAddRequest(req)
	created, err := f.uc.CreateFeedback(ctx, fb)
	if err != nil {
		f.log.Error("failed to create feedback", zap.Error(err))
		return nil, err
	}
	return &feedbacksvc.AddFeedbackResponse{
		Feedback: dto.ToProto(created),
	}, nil
}

func (f *Feedback) GetFeedback(ctx context.Context, req *feedbacksvc.GetFeedbackRequest) (*feedbacksvc.GetFeedbackResponse, error) {
	fb, err := f.uc.Get(ctx, req.FeedbackId)
	if err != nil {
		f.log.Error("failed to get feedback", zap.Error(err))
		return nil, err
	}
	return &feedbacksvc.GetFeedbackResponse{
		Feedback: dto.ToProto(fb),
	}, nil
}

func (f *Feedback) ListFeedbacks(ctx context.Context, req *feedbacksvc.ListCourseFeedbacksRequest) (*feedbacksvc.ListCourseFeedbacksResponse, error) {
	feedbacks, err := f.uc.GetCourseFeedbacks(ctx, req.CourseId)
	if err != nil {
		f.log.Error("failed to list feedbacks", zap.Error(err))
		return nil, err
	}

	resp := &feedbacksvc.ListCourseFeedbacksResponse{}
	for _, fb := range feedbacks {
		resp.Feedbacks = append(resp.Feedbacks, dto.ToProto(&fb))
	}
	resp.Total = int32(len(feedbacks))
	return resp, nil
}

func (f *Feedback) UpdateFeedback(ctx context.Context, req *feedbacksvc.UpdateFeedbackRequest) (*feedbacksvc.UpdateFeedbackResponse, error) {
	feedback := dto.ToModelFromUpdateRequest(req)
	err := f.uc.Update(ctx, req.FeedbackId, feedback)
	if err != nil {
		f.log.Error("failed to update feedback", zap.Error(err))
		return nil, err
	}
	return &feedbacksvc.UpdateFeedbackResponse{}, nil
}

func (f *Feedback) DeleteFeedback(ctx context.Context, req *feedbacksvc.DeleteFeedbackRequest) (*feedbacksvc.DeleteFeedbackResponse, error) {
	err := f.uc.Delete(ctx, req.FeedbackId)
	if err != nil {
		f.log.Error("failed to delete feedback", zap.Error(err))
		return nil, err
	}
	return &feedbacksvc.DeleteFeedbackResponse{}, nil
}

func (f *Feedback) GetAverageRating(ctx context.Context, req *feedbacksvc.GetAverageRatingRequest) (*feedbacksvc.GetAverageRatingResponse, error) {
	avg, err := f.uc.GetCourseRating(ctx, req.CourseId)
	if err != nil {
		f.log.Error("failed to get average rating", zap.Error(err))
		return nil, err
	}

	counted, err := f.uc.GetCourseFeedbacks(ctx, req.CourseId)
	if err != nil {
		return nil, err
	}

	return &feedbacksvc.GetAverageRatingResponse{
		AverageRating: float64(avg),
		FeedbackCount: int32(len(counted)),
	}, nil
}
