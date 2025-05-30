package usecase

import (
	"context"
	"testing"

	"github.com/sherinur/doit-platform/social-service/internal/model"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	CreateFn             func(ctx context.Context, userID int64, courseID, comment string, rating int32) (*model.Feedback, error)
	GetFn                func(ctx context.Context, id string) (*model.Feedback, error)
	GetCourseFeedbacksFn func(ctx context.Context, courseID string) ([]model.Feedback, error)
	UpdateFn             func(ctx context.Context, id string) error
	DeleteFn             func(ctx context.Context, id string) error
	ListFeedbacksFn      func(ctx context.Context) ([]model.Feedback, error)
}

func (m *mockRepo) ListFeedbacks(ctx context.Context) ([]model.Feedback, error) {
	return m.ListFeedbacksFn(ctx)
}

func (m *mockRepo) Create(ctx context.Context, userID int64, courseID, comment string, rating int32) (*model.Feedback, error) {
	return m.CreateFn(ctx, userID, courseID, comment, rating)
}

func (m *mockRepo) Get(ctx context.Context, id string) (*model.Feedback, error) {
	return m.GetFn(ctx, id)
}

func (m *mockRepo) GetCourseFeedbacks(ctx context.Context, courseID string) ([]model.Feedback, error) {
	return m.GetCourseFeedbacksFn(ctx, courseID)
}

func (m *mockRepo) Update(ctx context.Context, id string) error {
	return m.UpdateFn(ctx, id)
}

func (m *mockRepo) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}

type mockCache struct {
	store map[string]model.Feedback
}

func (m *mockCache) Set(fb model.Feedback) {
	m.store[fb.ID] = fb
}

func (m *mockCache) SetMany(fbs []model.Feedback) {
	for _, fb := range fbs {
		m.store[fb.ID] = fb
	}
}

func (m *mockCache) Get(id string) (model.Feedback, bool) {
	fb, ok := m.store[id]
	return fb, ok
}

func (m *mockCache) Delete(id string) {
	delete(m.store, id)
}

func newMockCache() *mockCache {
	return &mockCache{store: make(map[string]model.Feedback)}
}

func TestCreateFeedback(t *testing.T) {
	repo := &mockRepo{
		CreateFn: func(ctx context.Context, userID int64, courseID, comment string, rating int32) (*model.Feedback, error) {
			return &model.Feedback{
				ID:       "f1",
				UserID:   userID,
				CourseID: courseID,
				Comment:  comment,
				Rating:   rating,
			}, nil
		},
	}
	cache := newMockCache()
	uc := NewFeedback(repo, cache)

	fb := &model.Feedback{
		UserID:   1,
		CourseID: "c1",
		Comment:  "good",
		Rating:   5,
	}
	result, err := uc.CreateFeedback(context.Background(), fb)
	assert.NoError(t, err)
	assert.Equal(t, "f1", result.ID)
	_, ok := cache.Get("f1")
	assert.True(t, ok)
}

func TestCreateFeedback_Invalid(t *testing.T) {
	uc := NewFeedback(nil, nil)
	fb := &model.Feedback{}
	_, err := uc.CreateFeedback(context.Background(), fb)
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	repo := &mockRepo{
		GetFn: func(ctx context.Context, id string) (*model.Feedback, error) {
			return &model.Feedback{ID: id}, nil
		},
	}
	cache := newMockCache()
	uc := NewFeedback(repo, cache)

	fb, err := uc.Get(context.Background(), "f1")
	assert.NoError(t, err)
	assert.Equal(t, "f1", fb.ID)

	cached, ok := cache.Get("f1")
	assert.True(t, ok)
	assert.Equal(t, "f1", cached.ID)

	fb2, err := uc.Get(context.Background(), "f1")
	assert.NoError(t, err)
	assert.Equal(t, "f1", fb2.ID)
}

func TestGet_InvalidID(t *testing.T) {
	uc := NewFeedback(nil, nil)
	_, err := uc.Get(context.Background(), "")
	assert.ErrorIs(t, err, model.ErrInvalidID)
}

func TestDelete(t *testing.T) {
	var deletedID string
	repo := &mockRepo{
		DeleteFn: func(ctx context.Context, id string) error {
			deletedID = id
			return nil
		},
	}
	cache := newMockCache()
	cache.Set(model.Feedback{ID: "f1"})
	uc := NewFeedback(repo, cache)

	err := uc.Delete(context.Background(), "f1")
	assert.NoError(t, err)
	assert.Equal(t, "f1", deletedID)
	_, ok := cache.Get("f1")
	assert.False(t, ok)
}

func TestDelete_InvalidID(t *testing.T) {
	uc := NewFeedback(nil, nil)
	err := uc.Delete(context.Background(), "")
	assert.ErrorIs(t, err, model.ErrInvalidID)
}

func TestUpdate(t *testing.T) {
	var updatedID string
	repo := &mockRepo{
		UpdateFn: func(ctx context.Context, id string) error {
			updatedID = id
			return nil
		},
	}
	cache := newMockCache()
	cache.Set(model.Feedback{ID: "f1"})
	uc := NewFeedback(repo, cache)

	err := uc.Update(context.Background(), "f1", &model.Feedback{})
	assert.NoError(t, err)
	assert.Equal(t, "f1", updatedID)
	_, ok := cache.Get("f1")
	assert.False(t, ok)
}

func TestGetCourseRating(t *testing.T) {
	repo := &mockRepo{
		GetCourseFeedbacksFn: func(ctx context.Context, courseID string) ([]model.Feedback, error) {
			return []model.Feedback{{Rating: 4}, {Rating: 5}}, nil
		},
	}
	uc := NewFeedback(repo, nil)

	rating, err := uc.GetCourseRating(context.Background(), "c1")
	assert.NoError(t, err)
	assert.Equal(t, int32(4), rating)
}

func TestGetCourseRating_NoFeedbacks(t *testing.T) {
	repo := &mockRepo{
		GetCourseFeedbacksFn: func(ctx context.Context, courseID string) ([]model.Feedback, error) {
			return nil, nil
		},
	}
	uc := NewFeedback(repo, nil)

	rating, err := uc.GetCourseRating(context.Background(), "c1")
	assert.NoError(t, err)
	assert.Equal(t, int32(0), rating)
}

func TestGetCourseFeedbacks_InvalidCourseID(t *testing.T) {
	uc := NewFeedback(nil, nil)
	_, err := uc.GetCourseFeedbacks(context.Background(), "")
	assert.ErrorIs(t, err, model.ErrInvalidCourse)
}
