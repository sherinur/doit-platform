package model

type Course struct {
	ID           string
	Title        string
	Description  string
	CategoryID   string
	InstructorID string
	Tags         []string
	CreatedAt    int64
	UpdatedAt    int64
}
