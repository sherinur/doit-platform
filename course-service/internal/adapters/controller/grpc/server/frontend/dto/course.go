package dto

import (
	"github.com/sherinur/doit-platform/course-service/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseDTO struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Instructor  primitive.ObjectID   `json:"instructor"`
	Tags        []primitive.ObjectID `json:"tags"`
	Category    primitive.ObjectID   `json:"category"`
	CreatedAt   int64                `json:"created_at"`
}

func ToModel(dto *CourseDTO) *model.Course {
	tags := make([]string, len(dto.Tags))
	for i, tag := range dto.Tags {
		tags[i] = tag.Hex()
	}
	return &model.Course{
		ID:          dto.ID.Hex(),
		Title:       dto.Title,
		Description: dto.Description,
		Tags:        tags,
		CreatedAt:   dto.CreatedAt,
	}
}

func ToDTO(model *model.Course) *CourseDTO {
	id, _ := primitive.ObjectIDFromHex(model.ID)
	tags := make([]primitive.ObjectID, len(model.Tags))
	for i, tag := range model.Tags {
		objID, err := primitive.ObjectIDFromHex(tag)
		if err == nil {
			tags[i] = objID
		}
	}

	return &CourseDTO{
		ID:          id,
		Title:       model.Title,
		Description: model.Description,
		Tags:        tags,
		CreatedAt:   model.CreatedAt,
	}
}
