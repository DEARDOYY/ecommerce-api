package usecase

import (
	"context"
	"errors"
	"time"

	jwtPkg "ecommerce-api/pkg/jwt"

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

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, req RegisterRequest) (*domain.User, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
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

func (u *authUsecase) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	// 1. หา user จาก email
	user, err := u.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// 2. เช็ค password ตรงไหม
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// 3. สร้าง JWT token
	token, err := jwtPkg.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &LoginResponse{
		Token: token,
		User:  user,
	}, nil
}
