package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/course-service/internal/model"
)

type CourseRepo interface {
	Create(ctx context.Context, course *model.Course) error
	Update(ctx context.Context, id string, updated model.Course) error
	GetByID(ctx context.Context, id string) (*model.Course, error)
	Search(ctx context.Context, filter string, page, pageSize int64) ([]*model.Course, error)
}

type CourseUsecase struct {
	repo CourseRepo
}

func NewCourseUsecase(repo CourseRepo) *CourseUsecase {
	return &CourseUsecase{repo: repo}
}

func (u *CourseUsecase) CreateCourse(ctx context.Context, course *model.Course) error {
	return u.repo.Create(ctx, course)
}

func (u *CourseUsecase) UpdateCourse(ctx context.Context, id string, updated model.Course) error {
	return u.repo.Update(ctx, id, updated)
}

func (u *CourseUsecase) GetCourse(ctx context.Context, id string) (*model.Course, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *CourseUsecase) SearchCourses(ctx context.Context, filter string, page, pageSize int64) ([]*model.Course, error) {
	return u.repo.Search(ctx, filter, page, pageSize)
}
