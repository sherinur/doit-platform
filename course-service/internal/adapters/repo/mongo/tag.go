package mongoRepo

import (
	"context"

	"github.com/sherinur/doit-platform/course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepository struct {
	collection *mongo.Collection
}

func NewTagRepository(db *mongo.Database) *TagRepository {
	return &TagRepository{collection: db.Collection("tags")}
}

func (r *TagRepository) Create(ctx context.Context, tag *model.Tag) error {
	tag.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, tag)
	return err
}

func (r *TagRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.Tag, error) {
	var tag model.Tag
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}
