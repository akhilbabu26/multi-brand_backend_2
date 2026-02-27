// package routes

// import (
// 	"github.com/akhilbabu26/multi-brand_backend_2/internal/auth"

// 	"github.com/gin-gonic/gin"
// )

// func Setup() *gin.Engine {

// 	r := gin.Default()

// 	authRoute := r.Group("/auth")
// 	{
// 		authRoute.POST("/signup", auth.Signup)
// 		authRoute.POST("/verify-otp", auth.VerifyOTP)
// 		authRoute.POST("/login", auth.Login)
// 		authRoute.POST("/refresh", auth.RefreshToken)
// 	}

// 	return r
// }

package routes

import (
	"github.com/akhilbabu26/multi-brand_backend_2/internal/auth"

	"github.com/gin-gonic/gin"
)

//
// ======================================================
// ROUTER SETUP
// ======================================================
//

// Setup initializes and returns the main Gin router
func Setup() *gin.Engine {

	// create gin engine
	r := gin.Default()

	//
	// --------------------------------------------------
	// AUTH ROUTES
	// --------------------------------------------------
	//

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/signup", auth.Signup)
		authRoute.POST("/verify-otp", auth.VerifyOTP)
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/refresh", auth.RefreshToken)
	}

	return r
}