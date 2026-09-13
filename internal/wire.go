package internal

import (
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"
	"ecommerce-api/pkg/database"
	"os"
)

type Handlers struct {
	Auth *handler.AuthHandler
	User *handler.UserHandler
}

func InitHandlers() *Handlers {

	userCollection := database.GetCollection(os.Getenv("MONGO_DB_NAME"), "users")

	userRepo := repository.NewUserRepository(userCollection)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	return &Handlers{
		Auth: authHandler,
		User: userHandler,
	}
}
