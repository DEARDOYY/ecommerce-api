package usecase

import (
	"context"
	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/repository"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRequest struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
}

type CategoryUsecase interface {
	CreateCategory(req CategoryRequest) error
	GetCategoriesAll(ctx context.Context) ([]*domain.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*domain.Category, error)
	GetCategoryBySlug(ctx context.Context, slug string) (*domain.Category, error)
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: repo}
}

func generateSlug(name string) string {
	// Implement your slug generation logic here
	// For example, you can convert the name to lowercase and replace spaces with hyphens
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func (u *categoryUsecase) CreateCategory(req CategoryRequest) error {

	category := &domain.Category{
		Name:      req.Name,
		Slug:      generateSlug(req.Name),
		CreatedAt: time.Now(),
	}
	return u.repo.Create(context.Background(), category)
}

func (u *categoryUsecase) GetCategoriesAll(ctx context.Context) ([]*domain.Category, error) {
	categories, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (u *categoryUsecase) GetCategoryByID(ctx context.Context, id string) (*domain.Category, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	category, err := u.repo.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (u *categoryUsecase) GetCategoryBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	category, err := u.repo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	return category, nil
}
