package dto

import (
	frontendv1 "github.com/sherinur/doit-platform/apis/gen/base/frontend/v1"
	svc "github.com/sherinur/doit-platform/apis/gen/content-service/service/frontend/file/v1"
	"github.com/sherinur/doit-platform/content-service/internal/model"
)

func ToFileFromCreateRequest(req *svc.CreateFileRequest) (model.File, error) {
	return model.File{
		Body: req.Body,
		Type: req.Type,
		Size: int64(len(req.Body)),
	}, nil
}

func FromFile(file model.File) *frontendv1.File {
	return &frontendv1.File{
		Body: file.Body,
		Type: file.Type,
		Size: file.Size,
	}
}
