package response

import "github.com/gin-gonic/gin"

const (
	ErrCodeInvalidRequest = "INVALID_REQUEST"
	ErrCodeValidation     = "VALIDATION_ERROR"
	ErrCodeUnauthorized   = "UNAUTHORIZED"
	ErrCodeConflict       = "CONFLICT"
	ErrCodeInternal       = "INTERNAL_ERROR"
	ErrCodeNotFound       = "NOT_FOUND"
	ErrCodeForbidden      = "FORBIDDEN"
)

func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, gin.H{
		"error": gin.H{
			"code":    code,
			"message": message,
		},
	})
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{"data": data})
}

func Message(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}
