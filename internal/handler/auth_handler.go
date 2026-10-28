package handler

import (
	"net/http"

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
	if !bindJSON(c, &req) {
		return
	}

	user, err := h.usecase.CreateUser(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req usecase.LoginRequest

	if !bindJSON(c, &req) {
		return
	}

	result, err := h.usecase.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, response.ErrCodeUnauthorized, "Invalid email or password")
		return
	}

	response.Success(c, http.StatusOK, result)

}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req usecase.RefreshRequest
	if !bindJSON(c, &req) {
		return
	}

	pair, err := h.usecase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, response.ErrCodeUnauthorized, "Invalid or expired refresh token")
		return
	}

	response.Success(c, http.StatusOK, pair)
}
