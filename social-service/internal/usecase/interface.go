package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/social-service/internal/model"
)

type FeedbackRepo interface {
	Create(ctx context.Context, userId int64, courseId string, comment string, rating int32) (*model.Feedback, error)
	GetCourseFeedbacks(ctx context.Context, courseId string) ([]model.Feedback, error)
	Get(ctx context.Context, feedbackId string) (*model.Feedback, error)
	Update(ctx context.Context, feedbackId string) error
	Delete(ctx context.Context, feedbackId string) error
	ListFeedbacks(ctx context.Context) ([]model.Feedback, error)
}

type FeedbackCache interface {
	Set(feedback model.Feedback)
	SetMany(feedbacks []model.Feedback)
	Get(feedbackID string) (model.Feedback, bool)
	Delete(feedbackID string)
}
