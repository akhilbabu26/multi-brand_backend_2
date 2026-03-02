package products

import (
	"net/http"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/models"

	"github.com/gin-gonic/gin"
)

// ADMIN: CREATE PRODUCT
func CreateProduct(c *gin.Context){
	var body CreateProductDOT

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	product := models.Product{
		Name:               body.Name,
		Type:               body.Type,
		Color:              body.Color,
		CostPrice:          body.CostPrice,
		OriginalPrice:      body.OriginalPrice,
		DiscountPercentage: body.DiscountPercentage,
		ImageURL:           body.ImageURL,
		Description:        body.Description,
		Stock:              body.Stock,
		IsActive:           body.IsActive,
	}

	if err := config.DB.Create(&product).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create product",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "product created",
		"product": product,
	})
}

// ADMIN - UPDATE PRODUCT
func UpdateProduct(c *gin.Context){
	id := c.Param("id")

	var body UpdateProductDTO // updating feilds
	var product models.Product // all product

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	if err := config.DB.First(&product, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	// update only provided fields
	if body.Name != nil {
		product.Name = *body.Name
	}
	if body.Type != nil {
		product.Type = *body.Type
	}
	if body.Color != nil {
		product.Color = *body.Color
	}
	if body.CostPrice != nil {
		product.CostPrice = *body.CostPrice
	}
	if body.OriginalPrice != nil {
		product.OriginalPrice = *body.OriginalPrice
	}
	if body.DiscountPercentage != nil {
		product.DiscountPercentage = *body.DiscountPercentage
	}
	if body.ImageURL != nil {
		product.ImageURL = *body.ImageURL
	}
	if body.Description != nil {
		product.Description = *body.Description
	}
	if body.Stock != nil {
		product.Stock = *body.Stock
	}
	if body.IsActive != nil {
		product.IsActive = *body.IsActive
	}

	//Save triggers BeforeUpdate hook
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product updated",
		"product": product,
	})
}

// ADMIN - DELETE PRODUCT
func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "delete failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})
}

// PUBLIC: GET ALL PRODUCTS
func GetProducts(c *gin.Context){
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

// PUBLIC - GET SINGLE PRODUCT
func GetProductByID(c *gin.Context){
	id := c.Param("id")

	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

