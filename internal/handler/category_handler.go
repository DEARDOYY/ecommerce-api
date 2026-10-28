package handler

import (
	"ecommerce-api/internal/usecase"
	"ecommerce-api/pkg/response"

	"github.com/gin-gonic/gin"

	"ecommerce-api/internal/repository"
	"errors"
	"net/http"
)

type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryHandler(usecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		usecase: usecase,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req usecase.CategoryRequest
	if !bindJSON(c, &req) {
		return
	}

	if err := h.usecase.CreateCategory(req); err != nil {
		if errors.Is(err, repository.ErrCategoryExists) {
			response.Error(c, http.StatusConflict, response.ErrCodeConflict, "category already exists")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Message(c, http.StatusCreated, "Category created successfully")
}

func (h *CategoryHandler) GetCategoriesAll(c *gin.Context) {
	categories, err := h.usecase.GetCategoriesAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, categories)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := h.usecase.GetCategoryByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCategoryID) {
			response.Error(c, http.StatusBadRequest, response.ErrCodeInvalidRequest, "invalid category id")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, category)
}

func (h *CategoryHandler) GetCategoryBySlug(c *gin.Context) {
	slug := c.Param("slug")
	category, err := h.usecase.GetCategoryBySlug(c.Request.Context(), slug)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, category)
}
