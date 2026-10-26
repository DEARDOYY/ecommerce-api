package routes

import (
	"ecommerce-api/internal/handler"

	"ecommerce-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup, h *handler.UserHandler) {
	user := rg.Group("/users")
	{
		user.GET("", middleware.AuthRequired(), middleware.RequireRole("admin"), h.GetUserAll)
		user.GET("/email", h.GetUserByEmail)
		user.GET("/:id", h.GetUserByID)
		user.PUT("/:id", h.UpdateUser)
		user.PUT("/:id/role", h.UpdateUserRole)
	}
}
