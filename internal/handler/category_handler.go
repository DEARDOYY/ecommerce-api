package handler

import (
	"ecommerce-api/internal/usecase"

	"github.com/gin-gonic/gin"

	"ecommerce-api/internal/repository"
	"errors"
	"io"
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
	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.Is(err, io.EOF) {
			respondError(c, http.StatusBadRequest, ErrCodeInvalidRequest, "request body is required")
			return
		}
		respondError(c, http.StatusBadRequest, ErrCodeValidation, validationMessage(err))
		return
	}

	if err := h.usecase.CreateCategory(req); err != nil {
		if errors.Is(err, repository.ErrCategoryExists) {
			respondError(c, http.StatusConflict, ErrCodeConflict, "category already exists")
			return
		}
		respondError(c, http.StatusInternalServerError, ErrCodeInternal, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}
