package handler

import (
	"ecommerce-api/internal/usecase"

	"ecommerce-api/pkg/response"

	"github.com/gin-gonic/gin"

	"ecommerce-api/internal/repository"

	"ecommerce-api/internal/domain"
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

	response.Message(c, http.StatusCreated, "Product created successfully")
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.usecase.GetAllProducts(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.usecase.GetProductByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			response.Error(c, http.StatusNotFound, response.ErrCodeNotFound, "Product not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, product)
}

func (h *ProductHandler) GetProductByName(c *gin.Context) {
	name := c.Query("name")
	product, err := h.usecase.GetProductByName(c.Request.Context(), name)
	if err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			response.Error(c, http.StatusNotFound, response.ErrCodeNotFound, "Product not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req usecase.UpdateProductRequest
	if !bindJSON(c, &req) {
		return
	}

	if err := h.usecase.UpdateProduct(c.Request.Context(), id, req); err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			response.Error(c, http.StatusNotFound, response.ErrCodeNotFound, "Product not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "Product updated successfully")
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	var f domain.ProductFilter
	if err := c.ShouldBindQuery(&f); err != nil {
		response.Error(c, http.StatusBadRequest, response.ErrCodeBadRequest, "Invalid query parameters")
		return
	}

	products, count, err := h.usecase.SearchProduct(c.Request.Context(), f)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.ErrCodeInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, map[string]interface{}{
		"products": products,
		"count":    count,
	})
}
