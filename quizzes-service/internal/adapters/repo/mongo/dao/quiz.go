package dao

import (
	"quizzes-service/internal/model"
	"time"
)

type Quiz struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	CreatedBy   string    `bson:"created_by"`
	Status      string    `bson:"status"`
	QuestionIDs []string  `bson:"question_ids"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func FromQuiz(quiz model.Quiz) Quiz {
	return Quiz{
		ID:          quiz.ID,
		Title:       quiz.Title,
		Description: quiz.Description,
		CreatedBy:   quiz.CreatedBy,
		Status:      quiz.Status,
		QuestionIDs: quiz.QuestionIDs,
		CreatedAt:   quiz.CreatedAt,
		UpdatedAt:   quiz.UpdatedAt,
	}
}

func ToQuiz(quiz Quiz) model.Quiz {
	return model.Quiz{
		ID:          quiz.ID,
		Title:       quiz.Title,
		Description: quiz.Description,
		CreatedBy:   quiz.CreatedBy,
		Status:      quiz.Status,
		QuestionIDs: quiz.QuestionIDs,
		CreatedAt:   quiz.CreatedAt,
		UpdatedAt:   quiz.UpdatedAt,
	}
}
