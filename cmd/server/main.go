package main

import (
	"log"

	"ecommerce-api/internal"
	"ecommerce-api/internal/routes"
	"ecommerce-api/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	client := database.ConnectDB()
	defer client.Disconnect(nil)

	h := internal.InitHandlers()

	r := gin.Default()
	routes.SetupRoutes(r, h)
	r.Run(":8080")

}
