package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/internal/model"
)

type FilePresenter interface {
	Create(ctx context.Context, file model.File) (string, error)
	Get(ctx context.Context, key string) (*model.File, error)
	Delete(ctx context.Context, key string) error
}
