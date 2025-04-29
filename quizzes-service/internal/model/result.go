package model

import "time"

type Result struct {
	ID       string
	UserID   string
	QuizID   string
	Score    float64
	PassedAt time.Time
}
