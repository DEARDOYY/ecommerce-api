package handler

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
