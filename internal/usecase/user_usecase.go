package usecase

import (
	"context"

	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	GetUserAll(ctx context.Context) ([]*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUserAll(ctx context.Context) ([]*domain.User, error) {
	users, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user, err := u.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	err := u.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
