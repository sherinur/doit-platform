package usecase

import (
	"context"

	"content-service/internal/model"
)

type File struct {
	repo FileRepo
}

func NewFile(fileRepo FileRepo) *File {
	return &File{
		repo: fileRepo,
	}
}

func (f *File) Create(ctx context.Context, file model.File) (string, error) {
	return f.repo.Create(ctx, file)
}

func (f *File) Get(ctx context.Context, key string) (*model.File, error) {
	return f.repo.Get(ctx, key)
}

func (f *File) Delete(ctx context.Context, key string) error {
	return f.repo.Delete(ctx, key)
}
