package handler

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	ErrCodeInvalidRequest = "INVALID_REQUEST"
	ErrCodeValidation     = "VALIDATION_ERROR"
	ErrCodeUnauthorized   = "UNAUTHORIZED"
	ErrCodeConflict       = "CONFLICT"
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

// bindJSON binds the request body into dst. On failure it writes the 400
// response itself and returns false, so callers just `return`.
func bindJSON(c *gin.Context, dst any) bool {
	if err := c.ShouldBindJSON(dst); err != nil {
		if errors.Is(err, io.EOF) {
			respondError(c, http.StatusBadRequest, ErrCodeInvalidRequest, "request body is required")
			return false
		}
		respondError(c, http.StatusBadRequest, ErrCodeValidation, validationMessage(err))
		return false
	}
	return true
}

func validationMessage(err error) string {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) || len(ve) == 0 {
		return "invalid request body"
	}

	field := strings.ToLower(ve[0].Field())
	switch ve[0].Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "min":
		return field + " is too short"
	case "max":
		return field + " is too long"
	default:
		return field + " is invalid"
	}
}

func respondSuccess(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{"data": data})
}

func respondMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}
