package handler

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"ecommerce-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// bindJSON binds the request body into dst. On failure it writes the 400
// response itself and returns false, so callers just `return`.
func bindJSON(c *gin.Context, dst any) bool {
	if err := c.ShouldBindJSON(dst); err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(c, http.StatusBadRequest, response.ErrCodeInvalidRequest, "request body is required")
			return false
		}
		response.Error(c, http.StatusBadRequest, response.ErrCodeValidation, validationMessage(err))
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

// authorizeSelfOrAdmin allows the request only when the logged-in user owns
// the target id or is an admin. On failure it writes the 403 itself and
// returns false, so callers just `return`.
func authorizeSelfOrAdmin(c *gin.Context, targetID string) bool {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	if role == "admin" || userID == targetID {
		return true
	}

	response.Error(c, http.StatusForbidden, response.ErrCodeForbidden, "you can only access your own account")
	return false
}
