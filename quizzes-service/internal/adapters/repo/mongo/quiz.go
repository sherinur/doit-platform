package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"quizzes-service/internal/adapters/repo/mongo/dao"
	"quizzes-service/internal/model"
)

type QuizRepository struct {
	conn       *mongo.Database
	collection string
}

func NewQuizRepository(conn *mongo.Database) *QuizRepository {
	return &QuizRepository{
		conn:       conn,
		collection: "quizzes",
	}
}

func (repo *QuizRepository) CreateQuiz(ctx context.Context, quiz model.Quiz) (model.Quiz, error) {
	Quiz := dao.FromQuiz(quiz)
	res, err := repo.conn.Collection(repo.collection).InsertOne(ctx, Quiz)
	if err != nil {
		return model.Quiz{}, fmt.Errorf("quiz with ID %d has not been created: %w", quiz.ID, err)
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()

	return model.Quiz{ID: insertedID}, nil
}

func (repo *QuizRepository) GetQuizById(ctx context.Context, id string) (model.Quiz, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	var quiz dao.Quiz
	err = repo.conn.Collection(repo.collection).FindOne(ctx, bson.M{"_id": objID}).Decode(&quiz)
	if err != nil {
		return model.Quiz{}, fmt.Errorf("quiz with ID %s has not been found: %w", id, err)
	}

	return dao.ToQuiz(quiz), nil
}

func (repo *QuizRepository) UpdateQuiz(ctx context.Context, quiz model.Quiz) error {
	objID, err := primitive.ObjectIDFromHex(quiz.ID)
	if err != nil {
		return fmt.Errorf("invalid quiz ID: %w", err)
	}

	updateFields := bson.M{}

	if quiz.Title != "" {
		updateFields["title"] = quiz.Title
	}
	if quiz.Description != "" {
		updateFields["description"] = quiz.Description
	}
	if quiz.CreatedBy != "" {
		updateFields["created_by"] = quiz.CreatedBy
	}
	if quiz.Status != "" {
		updateFields["status"] = quiz.Status
	}
	if quiz.TotalPoints != 0 {
		updateFields["total_points"] = quiz.TotalPoints
	}

	// If no fields to update (besides updated_at), optionally return error
	if len(updateFields) == 0 {
		return fmt.Errorf("no fields provided to update")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}

	result, err := repo.conn.Collection(repo.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update Quiz with ID %s: %w", quiz.ID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("quiz with ID %s not found", quiz.ID)
	}

	return nil
}

func (repo *QuizRepository) ChangeTotalPointsQuiz(ctx context.Context, id string, change float64) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid quiz ID: %w", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$inc": bson.M{
			"total_points": change,
		},
	}

	result, err := repo.conn.Collection(repo.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update Quiz with ID %s: %w", id, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("quiz with ID %s not found", id)
	}

	return nil
}
func (repo *QuizRepository) DeleteQuiz(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	result, err := repo.conn.Collection(repo.collection).DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("failed to delete Quiz with ID %d: %w", id, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("quiz with ID %d not found", id)
	}
	return nil
}
