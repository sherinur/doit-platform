package handler

import (
	"fmt"
	"net/http"

	"github.com/sherinur/doit-platform/content-service/internal/adapter/controller/http/server/handler/dto"

	"github.com/gin-gonic/gin"
)

type File struct {
	uc FileUsecase
}

func NewFile(uc FileUsecase) *File {
	return &File{
		uc: uc,
	}
}

func (f *File) Create(c *gin.Context) {
	file, err := dto.ToFileFromCreateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	key, err := f.uc.Create(c.Request.Context(), file)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File updated successfully", "key": key})
}

func (f *File) Get(c *gin.Context) {
	key := c.Param("key")

	file, err := f.uc.Get(c.Request.Context(), key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, file.Type, file.Body)
}

func (f *File) Delete(c *gin.Context) {
	key := c.Param("key")

	err := f.uc.Delete(c.Request.Context(), key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
