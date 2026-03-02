package admin

import (
	"net/http"
	"strings"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/models"

	"github.com/gin-gonic/gin"
)

// GET ALL USERS
func GetAllUsers(c *gin.Context){
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

//UPDATE USERS
func UpdateUser(c *gin.Context){
	id := c.Param("id")

	var body UpdateUserDTO
	if err := c.ShouldBindJSON(&body); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if body.Name != "" {
		user.Name = body.Name
	}

	if body.Email != "" {
		user.Email = body.Email
	}

	if body.Role != "" {
		user.Role = strings.ToLower(body.Role)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
		"user":    user,
	})
}

// BLOCK / UNBLOCK USER
func BlockUser(c *gin.Context){

	id := c.Param("id")

	var body BlockUserDTO
	if err := c.ShouldBindJSON(&body); err != nil {c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil { c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// prevent admin blocking himself
	adminID := c.GetUint("user_id")
	if user.ID == adminID{
		c.JSON(http.StatusBadRequest, gin.H{ "error": "cannot block yourself"})
		return
	}

	user.IsBlocked = body.IsBlocked

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": "failed to update block status"})
		return
	}

	msg := "user unblocked"
	if body.IsBlocked {
		msg = "user blocked"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}