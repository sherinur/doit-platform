package file

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/internal/model"
	svc "github.com/sherinur/doit-platform/apis/gen/content-service/service/frontend/file/v1"
)

type File struct {
	client svc.FileServiceClient
}

func NewFile(client svc.FileServiceClient) *File {
	return &File{
		client: client,
	}
}

func (f *File) Create(ctx context.Context, file model.File) (string, error) {
	return "", nil
}

func (f *File) Get(ctx context.Context, key string) (*model.File, error) {
	return nil, nil
}

func (f *File) Delete(ctx context.Context, key string) error {
	return nil
}
