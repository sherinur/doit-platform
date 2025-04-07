package dto

import (
	"order_service/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func FromCreateOrderRequest(c *gin.Context) (model.Order, error) {
	var request CreateOrderRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		return model.Order{}, err
	}

	return model.Order{Name: request.Name, Price: request.Price}, err
}
