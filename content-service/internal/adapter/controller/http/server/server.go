package server

import (
	"content-service/config"
	"content-service/internal/adapter/controller/http/server/handler"

	"github.com/gin-gonic/gin"
)

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer

	fileHandler *handler.File
}
