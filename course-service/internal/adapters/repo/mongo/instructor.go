package mongoRepo

import (
	"context"

	"github.com/sherinur/doit-platform/course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InstructorRepository struct {
	collection *mongo.Collection
}

func NewInstructorRepository(db *mongo.Database) *InstructorRepository {
	return &InstructorRepository{collection: db.Collection("instructors")}
}

func (r *InstructorRepository) Create(ctx context.Context, instructor *model.Instructor) error {
	instructor.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(ctx, instructor)
	return err
}

func (r *InstructorRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.Instructor, error) {
	var instructor model.Instructor
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&instructor)
	if err != nil {
		return nil, err
	}
	return &instructor, nil
}
