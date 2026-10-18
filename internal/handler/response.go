package handler

import "github.com/gin-gonic/gin"

const (
	ErrCodeInvalidRequest = "INVALID_REQUEST"
	ErrCodeValidation     = "VALIDATION_ERROR"
	ErrCodeInternal       = "INTERNAL_ERROR"
)

func respondError(c *gin.Context, status int, code, message string) {
	c.JSON(status, gin.H{
		"error": gin.H{
			"code":    code,
			"message": message,
		},
	})
}
