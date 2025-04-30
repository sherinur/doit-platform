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

type QuestionRepository struct {
	conn       *mongo.Database
	collection string
}

func NewQuestionRepository(conn *mongo.Database) *QuestionRepository {
	return &QuestionRepository{
		conn:       conn,
		collection: "questions",
	}
}

func (repo *QuestionRepository) CreateQuestion(ctx context.Context, question model.Question) (model.Question, error) {
	Question := dao.FromQuestion(question)
	res, err := repo.conn.Collection(repo.collection).InsertOne(ctx, Question)
	if err != nil {
		return model.Question{}, fmt.Errorf("question with ID %d has not been created: %w", question.ID, err)
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()

	return model.Question{ID: insertedID}, nil
}

func (repo *QuestionRepository) CreateQuestions(ctx context.Context, questions []model.Question) ([]model.Question, error) {
	var daoQuestions []interface{}
	for _, question := range questions {
		daoQuestions = append(daoQuestions, dao.FromQuestion(question))
	}

	res, err := repo.conn.Collection(repo.collection).InsertMany(ctx, daoQuestions)
	if err != nil {
		return nil, fmt.Errorf("questions have not been created: %w", err)
	}

	if len(res.InsertedIDs) != len(questions) {
		return nil, fmt.Errorf("number of inserted IDs does not match number of questions")
	}

	for i, id := range res.InsertedIDs {
		objectID, ok := id.(primitive.ObjectID)
		if !ok {
			return nil, fmt.Errorf("failed to cast inserted ID to ObjectID")
		}
		questions[i].ID = objectID.Hex()
	}

	return questions, nil
}

func (repo *QuestionRepository) GetQuestionById(ctx context.Context, id string) (model.Question, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Question{}, fmt.Errorf("error converting ObjectID: %w", err)
	}

	var question dao.Question
	err = repo.conn.Collection(repo.collection).FindOne(ctx, bson.M{"_id": objID}).Decode(&question)
	if err != nil {
		return model.Question{}, fmt.Errorf("question with ID %s has not been found: %w", id, err)
	}

	return dao.ToQuestion(question), nil
}

func (repo *QuestionRepository) GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error) {
	cursor, err := repo.conn.Collection(repo.collection).Find(ctx, bson.M{"quiz_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Questions: %w", err)
	}
	defer cursor.Close(ctx)

	var questions []dao.Question
	if err = cursor.All(ctx, &questions); err != nil {
		return nil, fmt.Errorf("failed to decode Questions: %w", err)
	}

	var result []model.Question
	for _, question := range questions {
		result = append(result, dao.ToQuestion(question))
	}

	return result, nil
}

func (repo *QuestionRepository) UpdateQuestion(ctx context.Context, question model.Question) error {
	objID, err := primitive.ObjectIDFromHex(question.ID)
	if err != nil {
		return fmt.Errorf("error converting ObjectID: %w", err)
	}

	updateFields := bson.M{}

	if question.Text != "" {
		updateFields["text"] = question.Text
	}
	if question.Points > 0 {
		updateFields["points"] = question.Points
	}
	if question.Type != "" {
		updateFields["type"] = question.Type
	}
	if question.QuizID != "" {
		updateFields["quiz_id"] = question.QuizID
	}

	if len(updateFields) == 0 {
		return fmt.Errorf("no fields provided to update")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}

	result, err := repo.conn.Collection(repo.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update Quiz with ID %s: %w", question.ID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("quiz with ID %s not found", question.ID)
	}

	return nil
}

func (repo *QuestionRepository) DeleteQuestion(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting ObjectID: %w", err)
	}

	result, err := repo.conn.Collection(repo.collection).DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("failed to delete Question with ID %d: %w", id, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("question with ID %d not found", id)
	}
	return nil
}
