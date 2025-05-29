package usecase

import (
	"context"
	"course-service/internal/model"
)

type InstructorRepo interface {
	Create(ctx context.Context, instructor *model.Instructor) error
	FindByID(ctx context.Context, id string) (*model.Instructor, error)
	FindAll(ctx context.Context) ([]*model.Instructor, error)
}

type InstructorUsecase struct {
	repo InstructorRepo
}

func NewInstructorUsecase(repo InstructorRepo) *InstructorUsecase {
	return &InstructorUsecase{repo: repo}
}

func (uc *InstructorUsecase) Create(ctx context.Context, instructor *model.Instructor) error {
	return uc.repo.Create(ctx, instructor)
}

func (uc *InstructorUsecase) GetByID(ctx context.Context, id string) (*model.Instructor, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *InstructorUsecase) List(ctx context.Context) ([]*model.Instructor, error) {
	return uc.repo.FindAll(ctx)
}
