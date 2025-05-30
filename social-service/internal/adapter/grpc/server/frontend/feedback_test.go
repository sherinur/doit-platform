package frontend

import (
	"context"
	"errors"
	"testing"

	feedbacksvc "github.com/sherinur/doit-platform/apis/gen/social-service/service/frontend/feedback/v1"
	"github.com/sherinur/doit-platform/social-service/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mockUsecase struct {
	CreateFeedbackFn     func(context.Context, *model.Feedback) (*model.Feedback, error)
	GetFn                func(context.Context, string) (*model.Feedback, error)
	GetCourseFeedbacksFn func(context.Context, string) ([]model.Feedback, error)
	UpdateFn             func(context.Context, string, *model.Feedback) error
	DeleteFn             func(context.Context, string) error
	GetCourseRatingFn    func(context.Context, string) (int32, error)
}

func (m *mockUsecase) CreateFeedback(ctx context.Context, fb *model.Feedback) (*model.Feedback, error) {
	return m.CreateFeedbackFn(ctx, fb)
}

func (m *mockUsecase) Get(ctx context.Context, id string) (*model.Feedback, error) {
	return m.GetFn(ctx, id)
}

func (m *mockUsecase) GetCourseFeedbacks(ctx context.Context, courseID string) ([]model.Feedback, error) {
	return m.GetCourseFeedbacksFn(ctx, courseID)
}

func (m *mockUsecase) Update(ctx context.Context, id string, fb *model.Feedback) error {
	return m.UpdateFn(ctx, id, fb)
}

func (m *mockUsecase) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}

func (m *mockUsecase) GetCourseRating(ctx context.Context, courseID string) (int32, error) {
	return m.GetCourseRatingFn(ctx, courseID)
}

func TestAddFeedback(t *testing.T) {
	logger := zaptest.NewLogger(t)
	sampleModel := &model.Feedback{
		ID:        "f1",
		UserID:    1,
		CourseID:  "c1",
		Comment:   "good",
		Rating:    5,
		CreatedAt: timestamppb.Now().AsTime(),
	}

	uc := &mockUsecase{
		CreateFeedbackFn: func(ctx context.Context, fb *model.Feedback) (*model.Feedback, error) {
			return sampleModel, nil
		},
	}
	service := NewFeedback(uc, logger)

	req := &feedbacksvc.AddFeedbackRequest{
		UserId:   1,
		CourseId: "c1",
		Comment:  "good",
		Rating:   5,
	}

	resp, err := service.AddFeedback(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "f1", resp.Feedback.Id)
	assert.Equal(t, int64(1), resp.Feedback.UserId)
	assert.Equal(t, "c1", resp.Feedback.CourseId)
	assert.Equal(t, "good", resp.Feedback.Comment)
	assert.Equal(t, int32(5), resp.Feedback.Rating)
}

func TestGetFeedback(t *testing.T) {
	logger := zaptest.NewLogger(t)
	sampleModel := &model.Feedback{
		ID:        "f1",
		UserID:    2,
		CourseID:  "c2",
		Comment:   "great",
		Rating:    6,
		CreatedAt: timestamppb.Now().AsTime(),
	}

	uc := &mockUsecase{
		GetFn: func(ctx context.Context, id string) (*model.Feedback, error) {
			return sampleModel, nil
		},
	}
	service := NewFeedback(uc, logger)

	req := &feedbacksvc.GetFeedbackRequest{FeedbackId: "f1"}
	resp, err := service.GetFeedback(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "f1", resp.Feedback.Id)
	assert.Equal(t, "c2", resp.Feedback.CourseId)
}

func TestListFeedbacks(t *testing.T) {
	logger := zaptest.NewLogger(t)
	feedbacks := []model.Feedback{
		{
			ID:        "f1",
			UserID:    3,
			CourseID:  "c3",
			Comment:   "ok",
			Rating:    4,
			CreatedAt: timestamppb.Now().AsTime(),
		},
		{
			ID:        "f2",
			UserID:    4,
			CourseID:  "c3",
			Comment:   "nice",
			Rating:    5,
			CreatedAt: timestamppb.Now().AsTime(),
		},
	}

	uc := &mockUsecase{
		GetCourseFeedbacksFn: func(ctx context.Context, courseID string) ([]model.Feedback, error) {
			return feedbacks, nil
		},
	}
	service := NewFeedback(uc, logger)

	req := &feedbacksvc.ListCourseFeedbacksRequest{CourseId: "c3"}
	resp, err := service.ListFeedbacks(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Feedbacks, 2)
	assert.Equal(t, int32(2), resp.Total)
}

func TestDeleteFeedback(t *testing.T) {
	logger := zaptest.NewLogger(t)
	uc := &mockUsecase{
		DeleteFn: func(ctx context.Context, id string) error {
			if id == "notfound" {
				return errors.New("not found")
			}
			return nil
		},
	}
	service := NewFeedback(uc, logger)

	t.Run("delete existing", func(t *testing.T) {
		req := &feedbacksvc.DeleteFeedbackRequest{FeedbackId: "f1"}
		resp, err := service.DeleteFeedback(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("delete non-existing", func(t *testing.T) {
		req := &feedbacksvc.DeleteFeedbackRequest{FeedbackId: "notfound"}
		_, err := service.DeleteFeedback(context.Background(), req)
		assert.Error(t, err)
	})
}

func TestGetAverageRating(t *testing.T) {
	logger := zaptest.NewLogger(t)
	uc := &mockUsecase{
		GetCourseRatingFn: func(ctx context.Context, courseID string) (int32, error) {
			return 4, nil
		},
		GetCourseFeedbacksFn: func(ctx context.Context, courseID string) ([]model.Feedback, error) {
			return []model.Feedback{{}, {}, {}}, nil
		},
	}
	service := NewFeedback(uc, logger)

	req := &feedbacksvc.GetAverageRatingRequest{CourseId: "c5"}
	resp, err := service.GetAverageRating(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, float64(4), resp.AverageRating)
	assert.Equal(t, int32(3), resp.FeedbackCount)
}

func TestUpdateFeedback(t *testing.T) {
	logger := zaptest.NewLogger(t)
	var capturedID string
	var capturedModel *model.Feedback

	uc := &mockUsecase{
		UpdateFn: func(ctx context.Context, id string, fb *model.Feedback) error {
			capturedID = id
			capturedModel = fb
			if id == "bad" {
				return errors.New("update error")
			}
			return nil
		},
	}
	service := NewFeedback(uc, logger)

	t.Run("successful update", func(t *testing.T) {
		req := &feedbacksvc.UpdateFeedbackRequest{
			FeedbackId: "f1",
			Comment:    "updated comment",
			Rating:     4,
		}
		_, err := service.UpdateFeedback(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, "f1", capturedID)
		assert.Equal(t, "updated comment", capturedModel.Comment)
		assert.Equal(t, int32(4), capturedModel.Rating)
	})

	t.Run("update failure", func(t *testing.T) {
		req := &feedbacksvc.UpdateFeedbackRequest{
			FeedbackId: "bad",
			Comment:    "fail",
			Rating:     1,
		}
		_, err := service.UpdateFeedback(context.Background(), req)
		assert.Error(t, err)
	})
}
