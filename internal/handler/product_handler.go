package handler

import (
	"ecommerce-api/internal/usecase"

	"ecommerce-api/pkg/response"

	"github.com/gin-gonic/gin"

	"ecommerce-api/internal/repository"

	"errors"
	"net/http"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req usecase.ProductRequest
	if !bindJSON(c, &req) {
		return
	}

	if err := h.usecase.CreateProduct(req); err != nil {
		if err == usecase.ErrCategoryNotFound {
			response.Error(c, 404, response.ErrCodeNotFound, "Category not found")
			return
		}

		if errors.Is(err, repository.ErrCategoryExists) {
			response.Error(c, http.StatusConflict, response.ErrCodeConflict, "category already exists")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return

	}
}
