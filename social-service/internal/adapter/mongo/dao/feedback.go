package dao

import (
	"time"

	"github.com/sherinur/doit-platform/social-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feedback struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    int64              `bson:"user_id"`
	CourseID  string             `bson:"course_id"`
	Comment   string             `bson:"comment"`
	Rating    int32              `bson:"rating"`
	CreatedAt time.Time          `bson:"created_at"`
}

func (f *Feedback) ToModel() *model.Feedback {
	return &model.Feedback{
		ID:        f.ID.Hex(),
		UserID:    f.UserID,
		CourseID:  f.CourseID,
		Comment:   f.Comment,
		Rating:    f.Rating,
		CreatedAt: f.CreatedAt,
	}
}

func FromModel(m *model.Feedback) (*Feedback, error) {
	objID, err := primitive.ObjectIDFromHex(m.ID)
	if err != nil {
		return nil, err
	}

	return &Feedback{
		ID:        objID,
		UserID:    m.UserID,
		CourseID:  m.CourseID,
		Comment:   m.Comment,
		Rating:    m.Rating,
		CreatedAt: m.CreatedAt,
	}, nil
}
