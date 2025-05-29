package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DAO struct {
	CourseCollection     *mongo.Collection
	CategoryCollection   *mongo.Collection
	TagCollection        *mongo.Collection
	InstructorCollection *mongo.Collection
}

func NewDAO(db *mongo.Database) *DAO {
	return &DAO{
		CourseCollection:     db.Collection("courses"),
		CategoryCollection:   db.Collection("categories"),
		TagCollection:        db.Collection("tags"),
		InstructorCollection: db.Collection("instructors"),
	}
}
