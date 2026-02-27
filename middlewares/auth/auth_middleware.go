// package middleware

// import (
// 	"strings"
// 	"net/http"

// 	"github.com/akhilbabu26/multi-brand_backend_2/config"
// 	"github.com/akhilbabu26/multi-brand_backend_2/utils"

// 	"github.com/gin-gonic/gin"
// )

// func AuthMiddleware() gin.HandlerFunc{
// 	return func(c *gin.Context){
// 		auth := c.GetHeader("Authorization")

// 		if auth == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
// 			return
// 		}

// 		if !strings.HasPrefix(auth, "Bearer ") { //Validate Bearer format
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "invalid authorization header",
// 			})
// 			return
// 		}

// 		token := strings.TrimPrefix(auth, "Bearer ")
// 		token = strings.TrimSpace(token)

// 		claims, err := utils.ValidateToken(
// 			token,
// 			config.AppConfig.JWT.AccessSecretKey,
// 		)

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
// 			return
// 		}

// 		c.Set("user_id", claims.UserID)
// 		c.Set("role", claims.Role)

// 		c.Next()
// 	}
// }

package middleware

import (
	"net/http"
	"strings"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/utils"

	"github.com/gin-gonic/gin"
)

//
// ======================================================
// AUTH MIDDLEWARE
// ======================================================
//

// AuthMiddleware validates JWT access token
// and attaches user data to request context.
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		//
		// --------------------------------------------------
		// READ AUTH HEADER
		// --------------------------------------------------
		//

		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			return
		}

		//
		// --------------------------------------------------
		// VALIDATE BEARER FORMAT
		// --------------------------------------------------
		//

		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			return
		}

		//
		// --------------------------------------------------
		// EXTRACT TOKEN
		// --------------------------------------------------
		//

		token := strings.TrimPrefix(auth, "Bearer ")
		token = strings.TrimSpace(token)

		//
		// --------------------------------------------------
		// VALIDATE JWT
		// --------------------------------------------------
		//

		claims, err := utils.ValidateToken(
			token,
			config.AppConfig.JWT.AccessSecretKey,
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		//
		// --------------------------------------------------
		// SET CONTEXT VALUES
		// --------------------------------------------------
		//

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}