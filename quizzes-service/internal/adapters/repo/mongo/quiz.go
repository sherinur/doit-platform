package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"quizzes-service/internal/adapters/repo/mongo/dao"
	"quizzes-service/internal/model"
	"time"
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

func (repo *QuizRepository) GetQuizAll(ctx context.Context) ([]model.Quiz, error) {
	cursor, err := repo.conn.Collection(repo.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Quizzes: %w", err)
	}
	defer cursor.Close(ctx)

	var quizzes []dao.Quiz
	if err = cursor.All(ctx, &quizzes); err != nil {
		return nil, fmt.Errorf("failed to decode Quizzes: %w", err)
	}

	var result []model.Quiz
	for _, quiz := range quizzes {
		result = append(result, dao.ToQuiz(quiz))
	}

	return result, nil
}

func (repo *QuizRepository) UpdateQuiz(ctx context.Context, quiz model.Quiz) error {
	objID, err := primitive.ObjectIDFromHex(quiz.ID)
	if err != nil {
		panic(err)
	}

	update := bson.M{
		"$set": bson.M{
			"title":        quiz.Title,
			"description":  quiz.Description,
			"created_by":   quiz.CreatedBy,
			"status":       quiz.Status,
			"question_ids": quiz.QuestionIDs,
			"updated_at":   time.Now(),
		},
	}

	filter := bson.M{"_id": objID}
	result, err := repo.conn.Collection(repo.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update Quiz with ID %d: %w", quiz.ID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("quiz with ID %d not found", quiz.ID)
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
