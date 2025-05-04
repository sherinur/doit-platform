package server

import (
	"fmt"

	"github.com/sherinur/doit-platform/content-service/config"
	"github.com/sherinur/doit-platform/content-service/internal/adapter/controller/http/server/handler"

	"github.com/gin-gonic/gin"
)

const serverIPAddress = "0.0.0.0:%d"

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string

	fileHandler *handler.File
}

func New(cfg config.Config, fileUsecase FileUsecase) *API {
	gin.SetMode(cfg.Server.HTTPServer.Mode)
	server := gin.New()
	server.Use(gin.Recovery())

	fileHandler := handler.NewFile(fileUsecase)

	api := &API{
		server:      server,
		cfg:         cfg.Server.HTTPServer,
		addr:        fmt.Sprintf(serverIPAddress, cfg.Server.HTTPServer.Port),
		fileHandler: fileHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	v1 := a.server.Group("/api/v1")
	{
		clients := v1.Group("/file")
		{
			clients.PUT("/", a.fileHandler.Create)
			clients.GET("/:key", a.fileHandler.Get)
			clients.DELETE("/:key", a.fileHandler.Delete)
		}
	}
}

func (a *API) Run() error {
	return a.server.Run(a.addr)
}
