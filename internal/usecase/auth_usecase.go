package usecase

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/repository"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=user admin"`
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, req RegisterRequest) (*domain.User, error)
}

type authUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &authUsecase{repo: repo}
}

func (u *authUsecase) CreateUser(ctx context.Context, req RegisterRequest) (*domain.User, error) {
	// Check if user already exists
	// existingUser, err := u.repo.FindByEmail(ctx, req.Email)
	// if err != nil && err != mongo.ErrNoDocuments {
	// 	return nil, err
	// }
	// if existingUser != nil {
	// 	return nil, errors.New("user already exists")
	// }

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	role := req.Role
	if role == "" {
		role = "user"
	}

	user := &domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
