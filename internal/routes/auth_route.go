package routes

import (
	"ecommerce-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoute(rg *gin.RouterGroup, h *handler.AuthHandler) {
	user := rg.Group("/auth")
	{
		user.POST("/register", h.Register)
	}
}
