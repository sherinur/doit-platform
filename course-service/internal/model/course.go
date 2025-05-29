package model

type Course struct {
	ID           string   `bson:"_id,omitempty"`
	Title        string   `bson:"title"`
	Description  string   `bson:"description"`
	CategoryID   string   `bson:"category_id"`
	InstructorID string   `bson:"instructor_id"`
	Tags         []string `bson:"tags"`
	CreatedAt    int64    `bson:"created_at"`
	UpdatedAt    int64    `bson:"updated_at"`
}
