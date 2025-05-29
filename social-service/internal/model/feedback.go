package model

import (
	"errors"
	"time"
)

var (
	ErrInvalidRating = errors.New("rating must be between 1 and 6")
	ErrEmptyComment  = errors.New("comment cannot be empty")
	ErrInvalidUserID = errors.New("user ID must be positive")
	ErrInvalidCourse = errors.New("course ID cannot be empty")
	ErrInvalidID     = errors.New("ID must be a non-empty string")
)

type Feedback struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    int64     `json:"user_id" bson:"user_id"`
	CourseID  string    `json:"course_id" bson:"course_id"`
	Comment   string    `json:"comment" bson:"comment"`
	Rating    int32     `json:"rating" bson:"rating"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func (f *Feedback) Validate() error {
	switch {
	case f.ID == "":
		return ErrInvalidID
	case f.Rating < 1 || f.Rating > 6:
		return ErrInvalidRating
	case f.Comment == "":
		return ErrEmptyComment
	case f.UserID <= 0:
		return ErrInvalidUserID
	case f.CourseID == "":
		return ErrInvalidCourse
	}
	return nil
}
