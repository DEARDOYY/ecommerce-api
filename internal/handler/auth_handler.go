package handler

import (
	"ecommerce-api/internal/usecase"
	"ecommerce-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(usecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req usecase.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	user, err := h.usecase.CreateUser(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, user)
}
