package routes

import (

	"github.com/akhilbabu26/multi-brand_backend_2/internal/auth"
	"github.com/akhilbabu26/multi-brand_backend_2/middlewares/auth_middleware"

	"github.com/gin-gonic/gin"
)
func Setup() *gin.Engine {

	r := gin.Default()

	// AUTH ROUTES (PUBLIC)
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/signup", auth.Signup)
		authRoute.POST("/verify-otp", auth.VerifyOTP)
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/refresh", auth.RefreshToken)
		authRoute.POST("/forgot-password", auth.ForgotPassword)
		authRoute.POST("/reset-password", auth.ResetPassword)
	}

	// USER ROUTES (LOGGED IN)
	userRoute := r.Group("/user")
	userRoute.Use(middleware.Authentication())
	{
		userRoute.GET("/profile", func(c *gin.Context){
			c.JSON(200, gin.H{"welcome": "user"})
		})
		// userRoute.PUT("/profile", user.UpdateProfile)
	}

	// ADMIN ROUTES
	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.Authentication("admin"))
	{
		adminRoute.GET("/dashboard", func(c *gin.Context){
			c.JSON(200, gin.H{"welcome": "admin dash board"})
		})
	}

	return r
}