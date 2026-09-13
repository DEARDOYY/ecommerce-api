package routes

import (
	"ecommerce-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup, h *handler.UserHandler) {
	user := rg.Group("/users")
	{
		user.GET("", h.GetUserAll)
		user.GET("/email", h.GetUserByEmail)
		user.GET("/:id", h.GetUserByID)
		user.PUT("/:id", h.UpdateUser)
	}
}
