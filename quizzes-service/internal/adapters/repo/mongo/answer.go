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

type AnswerRepository struct {
	conn       *mongo.Database
	collection string
}

func NewAnswerRepository(conn *mongo.Database) *AnswerRepository {
	return &AnswerRepository{
		conn:       conn,
		collection: "answers",
	}
}

func (repo *AnswerRepository) CreateAnswer(ctx context.Context, answer model.Answer) (model.Answer, error) {
	Answer := dao.FromAnswer(answer)
	res, err := repo.conn.Collection(repo.collection).InsertOne(ctx, Answer)
	if err != nil {
		return model.Answer{}, fmt.Errorf("answer with ID %d has not been created: %w", answer.ID, err)
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()

	return model.Answer{ID: insertedID}, nil
}

func (repo *AnswerRepository) GetAnswerById(ctx context.Context, id string) (model.Answer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Answer{}, fmt.Errorf("error converting ObjectID: %w", err)
	}

	var answer dao.Answer
	err = repo.conn.Collection(repo.collection).FindOne(ctx, bson.M{"_id": objID}).Decode(&answer)
	if err != nil {
		return model.Answer{}, fmt.Errorf("answer with ID %s has not been found: %w", id, err)
	}

	return dao.ToAnswer(answer), nil
}

func (repo *AnswerRepository) GetAnswersByQuestionId(ctx context.Context, id string) ([]model.Answer, error) {
	cursor, err := repo.conn.Collection(repo.collection).Find(ctx, bson.M{"question_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Answers: %w", err)
	}
	defer cursor.Close(ctx)

	var answers []dao.Answer
	if err = cursor.All(ctx, &answers); err != nil {
		return nil, fmt.Errorf("failed to decode Answers: %w", err)
	}

	var result []model.Answer
	for _, answer := range answers {
		result = append(result, dao.ToAnswer(answer))
	}

	return result, nil
}

func (repo *AnswerRepository) GetAnswersByQuestionIds(ctx context.Context, ids []string) ([]model.Answer, error) {
	filter := bson.M{
		"question_id": bson.M{"$in": ids},
	}

	cursor, err := repo.conn.Collection(repo.collection).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Answers: %w", err)
	}
	defer cursor.Close(ctx)

	var answers []dao.Answer
	if err = cursor.All(ctx, &answers); err != nil {
		return nil, fmt.Errorf("failed to decode Answers: %w", err)
	}

	var result []model.Answer
	for _, answer := range answers {
		result = append(result, dao.ToAnswer(answer))
	}

	return result, nil
}

func (repo *AnswerRepository) UpdateAnswer(ctx context.Context, answer model.Answer) error {
	objID, err := primitive.ObjectIDFromHex(answer.ID)
	if err != nil {
		return fmt.Errorf("error converting ObjectID: %w", err)
	}

	updateFields := bson.M{}

	if answer.Text != "" {
		updateFields["text"] = answer.Text
	}

	if answer.QuestionID != "" {
		updateFields["type"] = answer.QuestionID
	}

	if len(updateFields) == 0 {
		return fmt.Errorf("no fields provided to update")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}

	result, err := repo.conn.Collection(repo.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update Quiz with ID %s: %w", answer.ID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("quiz with ID %s not found", answer.ID)
	}

	return nil
}

func (repo *AnswerRepository) DeleteAnswer(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting ObjectID: %w", err)
	}

	result, err := repo.conn.Collection(repo.collection).DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("failed to delete Answer with ID %d: %w", id, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("answer with ID %d not found", id)
	}
	return nil
}
