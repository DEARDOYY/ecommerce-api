package routes

import (
	"ecommerce-api/internal/handler"

	"ecommerce-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup, h *handler.ProductHandler) {
	product := rg.Group("/products", middleware.AuthRequired())
	{
		product.POST("", middleware.RequireRole("admin"), h.CreateProduct)
	}
}
