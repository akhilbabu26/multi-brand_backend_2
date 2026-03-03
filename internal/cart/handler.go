package cart

import (
	"net/http"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/models"
	"github.com/gin-gonic/gin"
)

// add to cart
func AddToCart(c *gin.Context){

	userID := c.GetUint("user_id")

	var body AddCartDTO
	var cart models.Cart

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// checking cart has already that item
	err := config.DB.
		Where("user_id = ? AND product_id = ?", userID, body.ProductID).
		First(&cart).Error

	if err == nil{
		// already exists then increase qty
		cart.Quantity += body.Quantity
		config.DB.Save(&cart)
	}else{
		// createing new cart item
		cart = models.Cart{
			UserID: userID,
			ProductID: body.ProductID,
			Quantity: body.Quantity,
		}
		config.DB.Create(&cart)
	}

	c.JSON(http.StatusOK, gin.H{"message": "added to cart"})
}

// Get cart
func GetCart(c *gin.Context){
	userID := c.GetUint("user_id")

	var items []models.Cart

	config.DB.Where("user_id = ?", userID).Find(&items)

	c.JSON(http.StatusOK, gin.H{
		"cart": items,
	})
}

//updates the quantity of a cart item
func UpdateCart(c *gin.Context){
	userID := c.GetUint("user_id")
	cartID := c.Param("id")

	var body UpdateCartDTO
	var cart models.Cart

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := config.DB.
		Where("cart_id = ? AND user_id = ?", cartID, userID).
		First(&cart).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart item not found"})
		return
	}

	cart.Quantity = body.Quantity
	config.DB.Save(&cart)

	c.JSON(http.StatusOK, gin.H{"message": "cart updated"})
}

// remove from cart
func DeleteCart(c *gin.Context) {

	userID := c.GetUint("user_id")
	cartID := c.Param("id")

	err := config.DB.
		Where("cart_id = ? AND user_id = ?", cartID, userID).
		Delete(&models.Cart{}).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item removed"})
}