package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Timestamp string `json:"timestamp"`
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    Meta        `json:"meta"`
}

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    Meta        `json:"meta"`
}

func meta() Meta {
	return Meta{Timestamp: time.Now().UTC().Format(time.RFC3339)}
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta(),
	})
}

func Error(c *gin.Context, code int, message string, errs interface{}) {
	c.JSON(code, ErrorResponse{
		Status:  "error",
		Message: message,
		Errors:  errs,
		Meta:    meta(),
	})
}

// Shortcut ที่ใช้บ่อย
func OK(c *gin.Context, data interface{}) {
	Success(c, http.StatusOK, "Request successful", data)
}

func Created(c *gin.Context, data interface{}) {
	Success(c, http.StatusCreated, "Resource created successfully", data)
}

func BadRequest(c *gin.Context, message string, errs interface{}) {
	Error(c, http.StatusBadRequest, message, errs)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message, nil)
}

func Deleted(c *gin.Context, message string) {
	Success(c, http.StatusOK, message, nil)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, nil)
}
