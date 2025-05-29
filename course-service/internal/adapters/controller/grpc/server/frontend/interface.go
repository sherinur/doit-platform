package frontend

import (
	"context"
	"course-service/internal/model"
)

type CourseUsecase interface {
	CreateCourse(ctx context.Context, course *model.Course) error
	UpdateCourse(ctx context.Context, id string, updated model.Course) error
	GetCourse(ctx context.Context, id string) (*model.Course, error)
	SearchCourses(ctx context.Context, filter string, page, pageSize int64) ([]*model.Course, error)
}
