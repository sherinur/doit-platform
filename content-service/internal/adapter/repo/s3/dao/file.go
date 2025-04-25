package dao

import (
	"content-service/internal/model"
	"io"
)

type File struct {
	objectKey string
	body      io.Reader
}

func FromFile(file model.File) File {
	return File{
		objectKey: "name",
	}
}

func ToFile(file File) model.File {
	return model.File{}
}
