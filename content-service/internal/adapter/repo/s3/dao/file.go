package dao

import (
	"bytes"
	"io"

	"github.com/sherinur/doit-platform/content-service/internal/model"
)

type File struct {
	ObjectKey string
	Body      io.Reader
	Size      int64
	Type      string
}

func FromFile(file model.File) File {
	return File{
		ObjectKey: "file",
		Body:      bytes.NewReader(file.Body),
		Size:      file.Size,
		Type:      file.Type,
	}
}

func ToFile(data io.Reader, contentType string) model.File {
	body, _ := io.ReadAll(data)

	return model.File{
		Body: body,
		Size: int64(len(body)),
		Type: contentType,
	}
}
