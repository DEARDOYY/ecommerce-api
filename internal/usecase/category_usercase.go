package usecase

import (
	"context"
	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/repository"
)

type CategoryRequest struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
}

type CategoryUsecase interface {
	CreateCategory(req CategoryRequest) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: repo}
}

func (u *categoryUsecase) CreateCategory(req CategoryRequest) error {
	category := &domain.Category{
		Name: req.Name,
	}
	return u.repo.Create(context.Background(), category)
}
