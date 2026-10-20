package internal

import (
	"context"
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"
	"ecommerce-api/pkg/database"
	"log"
	"os"
	"time"
)

type Handlers struct {
	Auth     *handler.AuthHandler
	Category *handler.CategoryHandler
	User     *handler.UserHandler
}

func InitHandlers() *Handlers {

	userCollection := database.GetCollection(os.Getenv("MONGO_DB_NAME"), "users")
	userRepo := repository.NewUserRepository(userCollection)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	categoryCollection := database.GetCollection(os.Getenv("MONGO_DB_NAME"), "categories")
	categoryRepo := repository.NewCategoryRepository(categoryCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := categoryRepo.EnsureIndexes(ctx); err != nil {
		log.Fatalf("failed to create category indexes: %v", err)
	}

	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	return &Handlers{
		Auth:     authHandler,
		User:     userHandler,
		Category: categoryHandler,
	}
}
