package dao

import (
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
	"time"
)

type Quiz struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	Status      string    `json:"status"`
	TotalPoints float64   `json:"total_points"`
	CourseID    string    `json:"course_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromQuiz(quiz model.Quiz) Quiz {
	return Quiz{
		ID:          quiz.ID,
		Title:       quiz.Title,
		Description: quiz.Description,
		CreatedBy:   quiz.CreatedBy,
		Status:      quiz.Status,
		TotalPoints: quiz.TotalPoints,
		CourseID:    quiz.CourseID,
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
		TotalPoints: quiz.TotalPoints,
		CourseID:    quiz.CourseID,
		CreatedAt:   quiz.CreatedAt,
		UpdatedAt:   quiz.UpdatedAt,
	}
}
