package routes

import (
	"ecommerce-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(rg *gin.RouterGroup, h *handler.CategoryHandler) {
	category := rg.Group("/categories")
	{
		category.POST("", h.CreateCategory)
		category.GET("", h.GetCategoriesAll)

	}
}
