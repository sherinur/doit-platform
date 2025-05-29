package frontend

import (
	"context"

	"github.com/sherinur/doit-platform/social-service/internal/model"
)

type FeedbackUsecase interface {
	CreateFeedback(ctx context.Context, feedback *model.Feedback) (*model.Feedback, error)
	GetCourseFeedbacks(ctx context.Context, courseID string) ([]model.Feedback, error)
	Get(ctx context.Context, id string) (*model.Feedback, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, feedback *model.Feedback) error
	GetCourseRating(ctx context.Context, courseID string) (int32, error)
}
