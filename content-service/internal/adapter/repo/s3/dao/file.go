package dao

import (
	"bytes"
	"content-service/internal/model"
	"io"
)

type File struct {
	ObjectKey string
	Body      io.Reader
	Size      int64
}

func FromFile(file model.File) File {
	return File{
		ObjectKey: "file",
		Body:      bytes.NewReader(file.Body),
		Size:      file.Size,
	}
}

func ToFile(data io.Reader) model.File {
	body, _ := io.ReadAll(data)

	return model.File{
		Body: body,
		Size: int64(len(body)),
	}
}
