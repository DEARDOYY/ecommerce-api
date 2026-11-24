package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"ecommerce-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// bindJSON binds the request body into dst. On failure it writes the 400
// response itself and returns false, so callers just `return`.
func bindJSON(c *gin.Context, dst any) bool {
	if err := c.ShouldBindJSON(dst); err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(c, http.StatusBadRequest, response.ErrCodeInvalidRequest, "invalid request body")
			return false
		}

		var syntaxErr *json.SyntaxError
		var typeErr *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxErr), errors.Is(err, io.ErrUnexpectedEOF):
			response.Error(c, http.StatusBadRequest, response.ErrCodeInvalidRequest, "invalid request body")
			return false
		case errors.As(err, &typeErr):
			log.Printf("type error: field=%s expected=%s", typeErr.Field, typeErr.Type)
			response.Error(c, http.StatusBadRequest, response.ErrCodeInvalidRequest, "invalid request body")
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

	log.Printf("validation failed: field=%s tag=%s", ve[0].Field(), ve[0].Tag())
	return "invalid request"
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

	response.Error(c, http.StatusForbidden, response.ErrCodeForbidden, "forbidden")
	return false
}
