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

type ResultRepository struct {
	conn       *mongo.Database
	collection string
}

func NewResultRepository(conn *mongo.Database) *ResultRepository {
	return &ResultRepository{
		conn:       conn,
		collection: "results",
	}
}

func (repo *ResultRepository) CreateResult(ctx context.Context, result model.Result) (model.Result, error) {
	Result := dao.FromResult(result)
	res, err := repo.conn.Collection(repo.collection).InsertOne(ctx, Result)
	if err != nil {
		return model.Result{}, fmt.Errorf("result with ID %d has not been created: %w", result.ID, err)
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()

	return model.Result{ID: insertedID}, nil
}

func (repo *ResultRepository) GetResultById(ctx context.Context, id string) (model.Result, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Result{}, fmt.Errorf("error converting ObjectID: %w", err)
	}

	var result dao.Result
	err = repo.conn.Collection(repo.collection).FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return model.Result{}, fmt.Errorf("result with ID %s has not been found: %w", id, err)
	}

	return dao.ToResult(result), nil
}

func (repo *ResultRepository) GetResultsByQuizId(ctx context.Context, id string) ([]model.Result, error) {
	cursor, err := repo.conn.Collection(repo.collection).Find(ctx, bson.M{"quiz_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Results: %w", err)
	}
	defer cursor.Close(ctx)

	var results []dao.Result
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode Results: %w", err)
	}

	var output []model.Result
	for _, result := range results {
		output = append(output, dao.ToResult(result))
	}

	return output, nil
}

func (repo *ResultRepository) GetResultsByUserId(ctx context.Context, id string) ([]model.Result, error) {
	cursor, err := repo.conn.Collection(repo.collection).Find(ctx, bson.M{"user_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Results: %w", err)
	}
	defer cursor.Close(ctx)

	var results []dao.Result
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode Results: %w", err)
	}

	var output []model.Result
	for _, result := range results {
		output = append(output, dao.ToResult(result))
	}

	return output, nil
}

func (repo *ResultRepository) DeleteResult(ctx context.Context, id string) error {
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
