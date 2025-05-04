package dto

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/doit-platform/api-gateway/internal/model"
)

func ToFileFromCreateRequest(c *gin.Context) (model.File, error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return model.File{}, err
	}

	size := c.Request.ContentLength

	file := model.File{
		Body: body,
		Size: size,
		Type: c.GetHeader("Content-Type"),
	}

	return file, nil
}
