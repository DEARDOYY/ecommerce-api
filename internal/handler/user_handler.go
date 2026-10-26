package handler

import (
	"ecommerce-api/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetUserAll(c *gin.Context) {

	users, err := h.usecase.GetUserAll(c.Request.Context())
	if err != nil {
		respondError(c, http.StatusInternalServerError, ErrCodeInternal, err.Error())
		return
	}

	respondSuccess(c, http.StatusOK, users)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")

	user, err := h.usecase.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		respondError(c, http.StatusInternalServerError, ErrCodeInternal, err.Error())
		return
	}

	respondSuccess(c, http.StatusOK, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if !authorizeSelfOrAdmin(c, id) {
		return
	}

	user, err := h.usecase.GetUserByID(c.Request.Context(), id)
	if err != nil {
		respondError(c, http.StatusInternalServerError, ErrCodeInternal, err.Error())
		return
	}

	respondSuccess(c, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var req usecase.UpdateUserRequest
	if !bindJSON(c, &req) {
		return
	}

	if err := h.usecase.UpdateUser(c.Request.Context(), req, id); err != nil {
		respondError(c, http.StatusInternalServerError, ErrCodeInternal, err.Error())
		return
	}

	respondMessage(c, http.StatusOK, "User updated successfully")
}

func (h *UserHandler) UpdateUserRole(c *gin.Context) {
	id := c.Param("id")
	var req usecase.UpdateUserRoleRequest
	if !bindJSON(c, &req) {
		return
	}

	if err := h.usecase.UpdateUserRole(c.Request.Context(), id, req); err != nil {
		respondError(c, http.StatusInternalServerError, ErrCodeInternal, err.Error())
		return
	}

	respondMessage(c, http.StatusOK, "User role updated successfully")
}
