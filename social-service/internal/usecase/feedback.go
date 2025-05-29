package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/social-service/internal/model"
)

type Feedback struct {
	feedbackRepo FeedbackRepo
}

func NewFeedback(feedbackRepo FeedbackRepo) *Feedback {
	return &Feedback{
		feedbackRepo: feedbackRepo,
	}
}

func (f *Feedback) CreateFeedback(ctx context.Context, feedback *model.Feedback) (*model.Feedback, error) {
	if err := feedback.Validate(); err != nil {
		return nil, err
	}
	return f.feedbackRepo.Create(ctx, feedback.UserID, feedback.CourseID, feedback.Comment, feedback.Rating)
}

func (f *Feedback) GetCourseFeedbacks(ctx context.Context, courseID string) ([]model.Feedback, error) {
	if courseID == "" {
		return nil, model.ErrInvalidCourse
	}
	return f.feedbackRepo.GetCourseFeedbacks(ctx, courseID)
}

func (f *Feedback) Get(ctx context.Context, id string) (*model.Feedback, error) {
	if id == "" {
		return nil, model.ErrInvalidID
	}
	return f.feedbackRepo.Get(ctx, id)
}

func (f *Feedback) Delete(ctx context.Context, id string) error {
	if id == "" {
		return model.ErrInvalidID
	}
	return f.feedbackRepo.Delete(ctx, id)
}

func (f *Feedback) Update(ctx context.Context, id string, feedback *model.Feedback) error {
	// TODO
	return f.feedbackRepo.Update(ctx, id)
}

func (f *Feedback) GetCourseRating(ctx context.Context, courseID string) (int32, error) {
	feedbacks, err := f.feedbackRepo.GetCourseFeedbacks(ctx, courseID)
	if err != nil {
		return 0, err
	}

	if len(feedbacks) == 0 {
		return 0, nil
	}

	var sum int32
	for _, feedback := range feedbacks {
		sum += feedback.Rating
	}
	avg := sum / int32(len(feedbacks))

	return avg, nil
}
