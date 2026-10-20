package usecase

import (
	"context"
	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/repository"
	"strings"
	"time"
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
		Name:      req.Name,
		Slug:      generateSlug(req.Name),
		CreatedAt: time.Now(),
	}
	return u.repo.Create(context.Background(), category)
}

func generateSlug(name string) string {
	// Implement your slug generation logic here
	// For example, you can convert the name to lowercase and replace spaces with hyphens
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
