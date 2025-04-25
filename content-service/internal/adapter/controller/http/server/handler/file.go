package handler

import "github.com/gin-gonic/gin"

type File struct {
	uc FileUsecase
}

func NewFile(uc FileUsecase) *File {
	return &File{
		uc: uc,
	}
}

func (f *File) Create(c *gin.Context) {

}

func (f *File) Get(c *gin.Context) {

}

func (f *File) Delete(c *gin.Context) {

}
