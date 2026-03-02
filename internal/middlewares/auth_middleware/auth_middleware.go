package middleware

import (
	"net/http"
	"strings"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/utils"

	"github.com/gin-gonic/gin"
)


// it validates JWT. If roles are provided checks role also.
func Authentication(allowedRoles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {

		// AUTH HEADER CHECK
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			return
		}

		// PARSE AUTH HEADER (CASE INSENSITIVE)
		// Expected: "Bearer <token>"
		parts := strings.SplitN(auth, " ", 2)

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			return
		}

		token := strings.TrimSpace(parts[1])

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			return
		}

		// TOKEN VALIDATION
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

		// STORE USER DATA IN CONTEXT

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		// ROLE CHECK (OPTIONAL)

		if len(allowedRoles) > 0 {

			userRole := claims.Role
			allowed := false

			for _, role := range allowedRoles {
				if userRole == role {
					allowed = true
					break
				}
			}

			if !allowed {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error": "access denied",
				})
				return
			}
		}

		c.Next()
	}
}