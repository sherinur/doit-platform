package model

import "time"

type Quiz struct {
	ID          string
	Title       string
	Description string
	CreatedBy   string
	Status      string
	QuestionIDs []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
