package mongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) *Feedback {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	db := client.Database("test_feedback_db")
	err = db.Collection("feedbacks").Drop(context.Background())
	assert.NoError(t, err)

	return NewFeedback(db)
}

func TestMongoCreateAndGet(t *testing.T) {
	feedbackRepo := setupTestDB(t)

	fb, err := feedbackRepo.Create(context.Background(), 1, "course1", "test comment", 5)
	assert.NoError(t, err)
	assert.NotEmpty(t, fb.ID)
	assert.Equal(t, int64(1), fb.UserID)
	assert.Equal(t, "course1", fb.CourseID)
	assert.Equal(t, "test comment", fb.Comment)
	assert.Equal(t, int32(5), fb.Rating)

	fb2, err := feedbackRepo.Get(context.Background(), fb.ID)
	assert.NoError(t, err)
	assert.Equal(t, fb.ID, fb2.ID)
	assert.Equal(t, fb.UserID, fb2.UserID)
}

func TestMongoGetCourseFeedbacks(t *testing.T) {
	feedbackRepo := setupTestDB(t)

	_, err := feedbackRepo.Create(context.Background(), 1, "courseX", "comment1", 4)
	assert.NoError(t, err)
	_, err = feedbackRepo.Create(context.Background(), 2, "courseX", "comment2", 5)
	assert.NoError(t, err)

	fbs, err := feedbackRepo.GetCourseFeedbacks(context.Background(), "courseX")
	assert.NoError(t, err)
	assert.Len(t, fbs, 2)
}

func TestMongoDelete(t *testing.T) {
	feedbackRepo := setupTestDB(t)

	fb, err := feedbackRepo.Create(context.Background(), 1, "courseDel", "to be deleted", 3)
	assert.NoError(t, err)

	err = feedbackRepo.Delete(context.Background(), fb.ID)
	assert.NoError(t, err)

	_, err = feedbackRepo.Get(context.Background(), fb.ID)
	assert.Error(t, err)
}
