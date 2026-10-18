package handler

import (
	"ecommerce-api/internal/usecase"

	"github.com/gin-gonic/gin"
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateCategory(req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Category created successfully"})
}
