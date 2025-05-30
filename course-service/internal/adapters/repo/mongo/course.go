package mongoRepo

import (
	"context"
	"time"

	"github.com/sherinur/doit-platform/course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseRepository struct {
	collection *mongo.Collection
}

func NewCourseRepository(db *mongo.Database) *CourseRepository {
	return &CourseRepository{collection: db.Collection("courses")}
}

func (r *CourseRepository) Create(ctx context.Context, course *model.Course) error {
	course.ID = primitive.NewObjectID().Hex()
	now := time.Now().Unix()
	course.CreatedAt = now
	course.UpdatedAt = now
	_, err := r.collection.InsertOne(ctx, course)
	return err
}

func (r *CourseRepository) Update(ctx context.Context, id string, updated model.Course) error {
	updated.UpdatedAt = time.Now().Unix()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updated})
	return err
}

func (r *CourseRepository) GetByID(ctx context.Context, id string) (*model.Course, error) {
	var course model.Course
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&course)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) Search(ctx context.Context, filter string, page, pageSize int64) ([]*model.Course, error) {
	skip := (page - 1) * pageSize
	findOptions := options.Find().SetSkip(skip).SetLimit(pageSize)

	query := bson.M{}
	if filter != "" {
		query = bson.M{"title": bson.M{"$regex": filter, "$options": "i"}}
	}

	cursor, err := r.collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courses []*model.Course
	for cursor.Next(ctx) {
		var course model.Course
		if err := cursor.Decode(&course); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}
