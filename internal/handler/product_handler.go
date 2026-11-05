package handler

import (
	"ecommerce-api/internal/usecase"

	"github.com/gin-gonic/gin"
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
}
