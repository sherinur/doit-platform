package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/social-service/internal/model"
)

type Feedback struct {
	feedbackRepo  FeedbackRepo
	inmemoryCache FeedbackCache
}

func (f *Feedback) ListFeedbacks(context context.Context) (any, any) {
	panic("unimplemented")
}

func NewFeedback(feedbackRepo FeedbackRepo, cache FeedbackCache) *Feedback {
	return &Feedback{
		feedbackRepo:  feedbackRepo,
		inmemoryCache: cache,
	}
}

func (f *Feedback) CreateFeedback(ctx context.Context, feedback *model.Feedback) (*model.Feedback, error) {
	if err := feedback.Validate(); err != nil {
		return nil, err
	}
	created, err := f.feedbackRepo.Create(ctx, feedback.UserID, feedback.CourseID, feedback.Comment, feedback.Rating)
	if err != nil {
		return nil, err
	}
	f.inmemoryCache.Set(*created)
	return created, nil
}

func (f *Feedback) GetCourseFeedbacks(ctx context.Context, courseID string) ([]model.Feedback, error) {
	if courseID == "" {
		return nil, model.ErrInvalidCourse
	}
	feedbacks, err := f.feedbackRepo.GetCourseFeedbacks(ctx, courseID)
	if err != nil {
		return nil, err
	}
	f.inmemoryCache.SetMany(feedbacks)
	return feedbacks, nil
}

func (f *Feedback) Get(ctx context.Context, id string) (*model.Feedback, error) {
	if id == "" {
		return nil, model.ErrInvalidID
	}

	if cached, ok := f.inmemoryCache.Get(id); ok {
		return &cached, nil
	}

	fb, err := f.feedbackRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	f.inmemoryCache.Set(*fb)
	return fb, nil
}

func (f *Feedback) Delete(ctx context.Context, id string) error {
	if id == "" {
		return model.ErrInvalidID
	}
	err := f.feedbackRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	f.inmemoryCache.Delete(id)
	return nil
}

func (f *Feedback) Update(ctx context.Context, id string, feedback *model.Feedback) error {
	err := f.feedbackRepo.Update(ctx, id)
	if err != nil {
		return err
	}
	f.inmemoryCache.Delete(id)
	return nil
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
