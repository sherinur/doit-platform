package model

import "time"

type Quiz struct {
	ID          string
	Title       string
	Description string
	CreatedBy   string
	Status      string
	TotalPoints float64
	CourseID    string
	Questions   []Question
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
