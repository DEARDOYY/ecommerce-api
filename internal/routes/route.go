package routes

import (
	"ecommerce-api/internal"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *internal.Handlers) {
	api := r.Group("/api/v1")

	AuthRoute(api, h.Auth)

}
