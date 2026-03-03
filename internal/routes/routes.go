package routes

import (
	admin "github.com/akhilbabu26/multi-brand_backend_2/internal/admin_side"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/auth"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/cart"
	middleware "github.com/akhilbabu26/multi-brand_backend_2/internal/middlewares/auth_middleware"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/products"

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

	// PRODUCTS ROUTES (PUBLIC)
	r.GET("/products", products.GetProducts)
	r.GET("/products/:id", products.GetProductByID)

	// USER ROUTES (LOGGED IN)
	userRoute := r.Group("/user")
	userRoute.Use(middleware.Authentication())
	{
		userRoute.GET("/profile", func(c *gin.Context){c.JSON(200, gin.H{"welcome": "user"})})

		// CART ROUTES
		userRoute.POST("/cart", cart.AddToCart)
		userRoute.GET("/cart", cart.GetCart)
		userRoute.PUT("/cart/:id", cart.UpdateCart)
		userRoute.DELETE("/cart/:id", cart.DeleteCart)
	}

	// ADMIN ROUTES
	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.Authentication("admin"))
	{
		adminRoute.GET("/dashboard", func(c *gin.Context){c.JSON(200, gin.H{"welcome": "admin dash board"})})

		// user management
		adminRoute.GET("/users", admin.GetAllUsers)
		adminRoute.PUT("/users/:id", admin.UpdateUser)
		adminRoute.PUT("/users/:id/block", admin.BlockUser)

		// product management
		adminRoute.POST("/products", products.CreateProduct)
		adminRoute.PUT("/products/:id", products.UpdateProduct)
		adminRoute.DELETE("/products/:id", products.DeleteProduct)
	}

	return r
}