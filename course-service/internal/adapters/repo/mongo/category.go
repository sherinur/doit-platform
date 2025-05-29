package mongoRepo

import (
	"context"

	"github.com/sherinur/doit-platform/course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{collection: db.Collection("categories")}
}

func (r *CategoryRepository) Create(ctx context.Context, category *model.Category) error {
	category.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, category)
	return err
}

func (r *CategoryRepository) FindByID(ctx context.Context, id string) (*model.Category, error) {
	var category model.Category
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
