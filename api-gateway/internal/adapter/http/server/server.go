package server

import (
	"fmt"
	"net/http"

	"github.com/sherinur/doit-platform/api-gateway/config"
	"github.com/sherinur/doit-platform/api-gateway/internal/adapter/http/server/handler"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const serverIPAddress = "0.0.0.0:%d"

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string
	log    *zap.Logger

	fileHandler *handler.File
}

func New(cfg config.HTTPServer, logger *zap.Logger) *API {
	gin.SetMode(cfg.Mode)
	server := gin.New()
	server.Use(gin.Recovery())

	api := &API{
		server: server,
		cfg:    cfg,
		addr:   fmt.Sprintf(serverIPAddress, cfg.Port),

		log: logger,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	a.server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	v1 := a.server.Group("/api/v1")
	{
		file := v1.Group("/file")
		{
			file.PUT("/", a.fileHandler.Create)
			file.GET("/:key", a.fileHandler.Get)
			file.DELETE("/:key", a.fileHandler.Delete)
		}
	}
}

func (a *API) Run() error {
	a.log.Info("Running http server", zap.String("addr", a.addr))
	return a.server.Run(a.addr)
}
