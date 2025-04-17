package service

import (
	"inventory_service/config"
	"inventory_service/internal/adapter/http/service/handler"

	"github.com/gin-gonic/gin"
)

const serverIPAddress = "0.0.0.0:8081"

type API struct {
	server *gin.Engine
	cfg    *config.HTTPServer
	addr   string

	productHandler  *handler.Product
	discountHandler *handler.Discount
}

func New(cfg *config.Server, productHandler *handler.Product, discountHandler *handler.Discount) *API {
	gin.SetMode(cfg.HTTPServer.Mode)
	server := gin.New()

	server.Use(gin.Recovery())

	api := &API{
		server:         server,
		cfg:            &cfg.HTTPServer,
		addr:           serverIPAddress,
		productHandler: productHandler,
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

		// Products
		v1.POST("/products", a.productHandler.CreateProduct)
		v1.GET("/orders/:id", a.productHandler.GetProduct)
		v1.PATCH("/products/:id", a.productHandler.UpdateProduct)
		v1.DELETE("/products/:id", a.productHandler.DeleteProduct)
		v1.GET("/products", a.productHandler.ListProducts)

		// Promotion
		v1.POST("/discount", a.discountHandler.CreateDiscount)
		v1.GET("/promotion", a.discountHandler.GetAllProductsWithPromotion)
		v1.DELETE("/discount/:id", a.discountHandler.DeleteDiscount)
	}
}

func (api *API) Run() error {
	return api.server.Run(api.addr)
}

func (api *API) Stop() {}
