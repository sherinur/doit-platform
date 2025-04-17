package dto

import (
	"order_service/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required,min=1,dive"`
}

type OrderItemRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int   `json:"quantity" binding:"required,min=1"`
}

type OrderResponse struct {
	ID        int64               `json:"id"`
	Status    string              `json:"status"`
	Amount    float64             `json:"amount"`
	CreatedAt string              `json:"created_at"`
	Items     []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func FromCreateOrderRequest(c *gin.Context) (*CreateOrderRequest, error) {
	var request CreateOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func ToOrderResponse(order *model.Order) *OrderResponse {
	var items []OrderItemResponse
	for _, item := range order.OrderItems {
		items = append(items, OrderItemResponse{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.UnitPrice,
		})
	}

	return &OrderResponse{
		ID:        order.ID,
		Status:    order.Status,
		Amount:    order.TotalAmount,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		Items:     items,
	}
}

func (req *CreateOrderRequest) ToModel() *model.Order {
	items := make([]model.OrderItem, 0, len(req.Items))

	for _, item := range req.Items {
		items = append(items, model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: 0, // можно позже обновить через ProductService
		})
	}

	return &model.Order{
		Status:     "pending", // начальный статус
		OrderItems: items,
	}
}
