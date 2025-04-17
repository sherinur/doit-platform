package service

import (
	"order_service/config"
	"order_service/internal/adapter/http/service/handler"

	"github.com/gin-gonic/gin"
)

const serverIPAddress = "0.0.0.0:8080"

type API struct {
	server *gin.Engine
	cfg    *config.HTTPServer
	addr   string

	orderHandler *handler.Order
}

func New(cfg *config.Server, orderHandler *handler.Order) *API {
	gin.SetMode(cfg.HTTPServer.Mode)
	server := gin.New()

	server.Use(gin.Recovery())

	api := &API{
		server:       server,
		cfg:          &cfg.HTTPServer,
		addr:         serverIPAddress,
		orderHandler: orderHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	v1 := a.server.Group("api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "ok"})
		})

		v1.POST("/orders", a.orderHandler.CreateOrder)
		v1.GET("/orders/:id", a.orderHandler.GetOrder)
	}
}

func (api *API) Run() error {
	return api.server.Run(api.addr)
}

func (api *API) Stop() {}
