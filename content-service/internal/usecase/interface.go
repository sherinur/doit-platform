package usecase

import (
	"content-service/internal/model"
	"context"
)

type FileRepo interface {
	Create(ctx context.Context, file model.File) (string, error)
	Get(ctx context.Context, key string) (*model.File, error)
	Delete(ctx context.Context, key string) error
}
