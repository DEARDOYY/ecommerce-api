package middleware

import (
	jwtPkg "ecommerce-api/pkg/jwt"
	"ecommerce-api/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. ดึง token จาก header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header is required", nil)
			c.Abort()
			return
		}

		// 2. เช็ค format ต้องเป็น "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization format", nil)
			c.Abort()
			return
		}

		// 3. ถอดรหัส token
		claims, err := jwtPkg.ValidateToken(parts[1])
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid or expired token", nil)
			c.Abort()
			return
		}

		// 4. เก็บข้อมูล user ไว้ใน context ให้ handler เรียกใช้ได้
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}
