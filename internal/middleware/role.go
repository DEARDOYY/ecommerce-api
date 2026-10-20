package middleware

import (
	"ecommerce-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		userRole, _ := role.(string)

		for _, r := range roles {
			if userRole == r {
				c.Next()
				return
			}
		}

		response.Error(c, http.StatusForbidden, "Forbidden: insufficient permissions", nil)
		c.Abort()
	}
}
