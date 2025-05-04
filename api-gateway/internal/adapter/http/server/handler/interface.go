package handler

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/internal/model"
)

type FileUsecase interface {
	Create(ctx context.Context, file model.File) (string, error)
	Get(ctx context.Context, key string) (*model.File, error)
	Delete(ctx context.Context, key string) error
}
