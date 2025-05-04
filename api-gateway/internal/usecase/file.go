package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/internal/model"
)

type File struct {
	filePresenter FilePresenter
}

func NewFile(presenter FilePresenter) *File {
	return &File{
		filePresenter: presenter,
	}
}

func (f *File) Create(ctx context.Context, file model.File) (string, error) {
	return f.filePresenter.Create(ctx, file)
}

func (f *File) Get(ctx context.Context, key string) (*model.File, error) {
	return f.filePresenter.Get(ctx, key)
}

func (f *File) Delete(ctx context.Context, key string) error {
	return f.filePresenter.Delete(ctx, key)
}
