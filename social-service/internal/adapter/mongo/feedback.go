package mongo

import (
	"context"
	"time"

	"github.com/sherinur/doit-platform/social-service/internal/adapter/mongo/dao"
	"github.com/sherinur/doit-platform/social-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Feedback struct {
	conn       *mongo.Database
	collection string
}

const collectionFeedbacks = "feedbacks"

func NewFeedback(connection *mongo.Database) *Feedback {
	return &Feedback{
		conn:       connection,
		collection: collectionFeedbacks,
	}
}

func (f *Feedback) Create(ctx context.Context, userId int64, courseId string, comment string, rating int32) (*model.Feedback, error) {
	now := time.Now()
	daoFeedback := &dao.Feedback{
		ID:        primitive.NewObjectID(),
		UserID:    userId,
		CourseID:  courseId,
		Comment:   comment,
		Rating:    rating,
		CreatedAt: now,
	}

	_, err := f.conn.Collection(f.collection).InsertOne(ctx, daoFeedback)
	if err != nil {
		return nil, err
	}

	return daoFeedback.ToModel(), nil
}

func (f *Feedback) GetCourseFeedbacks(ctx context.Context, courseId string) ([]model.Feedback, error) {
	cursor, err := f.conn.Collection(f.collection).Find(ctx, bson.M{"course_id": courseId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var daoFeedbacks []dao.Feedback
	if err = cursor.All(ctx, &daoFeedbacks); err != nil {
		return nil, err
	}

	feedbacks := make([]model.Feedback, 0, len(daoFeedbacks))
	for _, df := range daoFeedbacks {
		feedbacks = append(feedbacks, *df.ToModel())
	}

	return feedbacks, nil
}

func (f *Feedback) Get(ctx context.Context, feedbackId string) (*model.Feedback, error) {
	objID, err := primitive.ObjectIDFromHex(feedbackId)
	if err != nil {
		return nil, err
	}

	var df dao.Feedback
	err = f.conn.Collection(f.collection).FindOne(ctx, bson.M{"_id": objID}).Decode(&df)
	if err != nil {
		return nil, err
	}

	return df.ToModel(), nil
}

func (f *Feedback) Update(ctx context.Context, feedbackId string) error {
	// TODO
	return nil
}

func (f *Feedback) Delete(ctx context.Context, feedbackId string) error {
	objID, err := primitive.ObjectIDFromHex(feedbackId)
	if err != nil {
		return err
	}

	_, err = f.conn.Collection(f.collection).DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (f *Feedback) ListFeedbacks(ctx context.Context) ([]model.Feedback, error) {
	cursor, err := f.conn.Collection(f.collection).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var daoFeedbacks []dao.Feedback
	if err = cursor.All(ctx, &daoFeedbacks); err != nil {
		return nil, err
	}

	feedbacks := make([]model.Feedback, 0, len(daoFeedbacks))
	for _, df := range daoFeedbacks {
		feedbacks = append(feedbacks, *df.ToModel())
	}

	return feedbacks, nil
}
