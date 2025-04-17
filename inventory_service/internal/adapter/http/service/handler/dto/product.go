package dto

import (
	"time"

	"inventory_service/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gte=0"`
	Category    string  `json:"category" binding:"required"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Category    *string  `json:"category,omitempty"`
}

type ProductResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func FromCreateProductRequest(c *gin.Context) (*CreateProductRequest, error) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func FromUpdateProductRequest(c *gin.Context) (*UpdateProductRequest, error) {
	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *CreateProductRequest) ToModel() *model.Product {
	return &model.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Category:    r.Category,
	}
}

func (r *UpdateProductRequest) ToModel() *model.Product {
	return &model.Product{
		Name:        *r.Name,
		Description: *r.Description,
		Price:       *r.Price,
		Category:    *r.Category,
	}
}

func ToProductResponse(p *model.Product) *ProductResponse {
	return &ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Category:    p.Category,
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
	}
}
