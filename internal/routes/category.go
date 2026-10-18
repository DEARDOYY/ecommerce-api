package routes

import (
	"ecommerce-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(rg *gin.RouterGroup, h *handler.CategoryHandler) {
	category := rg.Group("/categories")
	{
		category.POST("", h.CreateCategory)

		// category.GET("", h.GetCategories)
		// category.GET("/:id", h.GetCategoryByID)
		// category.POST("", h.CreateCategory)
		// category.PUT("/:id", h.UpdateCategory)
		// category.DELETE("/:id", h.DeleteCategory)
	}
}
