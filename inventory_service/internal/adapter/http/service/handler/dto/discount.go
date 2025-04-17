package dto

import (
	"net/http"
	"strconv"
	"time"

	"inventory_service/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateDiscountRequest struct {
	Name               string    `json:"name" binding:"required"`
	Description        string    `json:"description" binding:"required"`
	ProductId          int64     `json:"product_id"`
	DiscountPercentage float64   `json:"discount_percentage" binding:"required"`
	StartDate          time.Time `json:"start_date" binding:"required"`
	EndDate            time.Time `json:"end_date" binding:"required"`
	IsActive           bool      `json:"is_active" binding:"required"`
}

func (r *CreateDiscountRequest) ToModel() model.Discount {
	return model.Discount{
		Name:               r.Name,
		Description:        r.Description,
		DiscountPercentage: r.DiscountPercentage,
		StartDate:          r.StartDate,
		EndDate:            r.EndDate,
		IsActive:           r.IsActive,
	}
}

func FromCreateDiscountRequest(c *gin.Context) (*CreateDiscountRequest, error) {
	var req CreateDiscountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func ValidateDiscountID(c *gin.Context) (int64, error) {
	discountIDStr := c.Param("id")
	if discountIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Discount ID is required"})
		return 0, nil
	}

	discountID, err := strconv.ParseInt(discountIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Discount ID"})
		return 0, err
	}

	return discountID, nil
}
