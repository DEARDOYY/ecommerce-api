package routes

import (
	"ecommerce-api/internal/handler"

	"ecommerce-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(rg *gin.RouterGroup, h *handler.CategoryHandler) {
	category := rg.Group("/categories", middleware.AuthRequired())
	{
		category.POST("", middleware.RequireRole("admin"), h.CreateCategory)
		category.GET("", h.GetCategoriesAll)
		category.GET("/:id", h.GetCategoryByID)
		category.GET("/slug/:slug", h.GetCategoryBySlug)
	}
}
