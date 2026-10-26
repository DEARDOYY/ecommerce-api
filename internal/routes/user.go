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
		user.GET("/email", middleware.AuthRequired(), h.GetUserByEmail)
		user.GET("/:id", middleware.AuthRequired(), h.GetUserByID)
		user.PUT("/:id", middleware.AuthRequired(), h.UpdateUser)
		user.PUT("/:id/role", middleware.AuthRequired(), middleware.RequireRole("admin"), h.UpdateUserRole)
	}
}
