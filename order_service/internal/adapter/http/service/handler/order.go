package handler

import (
	"net/http"

	"order_service/internal/adapter/http/service/handler/dto"

	"github.com/gin-gonic/gin"
)

type Order struct {
	uc OrderUsecase
}

func NewOrder(uc OrderUsecase) *Order {
	return &Order{uc: uc}
}

func (o *Order) CreateOrder(c *gin.Context) {
	item, err := dto.FromCreateOrderRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = o.uc.Create(c.Request.Context(), item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
}
