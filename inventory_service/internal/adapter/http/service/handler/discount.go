package handler

import (
	"net/http"

	"inventory_service/internal/adapter/http/service/handler/dto"

	"github.com/gin-gonic/gin"
)

type Discount struct {
	uc DiscountUsecase
}

func NewDiscount(uc DiscountUsecase) *Discount {
	return &Discount{uc: uc}
}

func (d *Discount) CreateDiscount(c *gin.Context) {
	req, err := dto.FromCreateDiscountRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	discount := req.ToModel()

	err = d.uc.Create(c.Request.Context(), discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create discount"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Discount created successfully"})
}

func (d *Discount) GetAllProductsWithPromotion(c *gin.Context) {
	products, err := d.uc.GetAllProductsWithPromotion(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (d *Discount) DeleteDiscount(c *gin.Context) {
	discountID, err := dto.ValidateDiscountID(c)
	if err != nil {
		return
	}

	err = d.uc.Delete(c, discountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete discount"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Discount deleted successfully"})
}
