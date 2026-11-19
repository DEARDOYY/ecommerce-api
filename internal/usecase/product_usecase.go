package usecase

import (
	"context"
	"ecommerce-api/internal/repository"
	"errors"

	"ecommerce-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  string  `json:"category_id" binding:"required"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  string  `json:"category_id" binding:"required"`
}

type ProductUsecase interface {
	CreateProduct(req ProductRequest) error
	GetAllProducts(ctx context.Context) ([]*domain.Product, error)
	GetProductByID(ctx context.Context, id string) (*domain.Product, error)
	GetProductByName(ctx context.Context, name string) (*domain.Product, error)
}

var ErrCategoryNotFound = errors.New("category not found")
var ErrProductNotFound = errors.New("product not found")
var ErrInvalidProductID = errors.New("invalid product id")
var ErrInvalidProductName = errors.New("product name cannot be empty")

type productUsecase struct {
	repo         repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductUsecase(repo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductUsecase {
	return &productUsecase{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (u *productUsecase) CreateProduct(req ProductRequest) error {

	categoryID, err := primitive.ObjectIDFromHex(req.CategoryID)
	if err != nil {
		return ErrInvalidCategoryID
	}

	category, err := u.categoryRepo.FindByID(context.Background(), categoryID)
	if err != nil {
		return err
	}
	if category == nil {
		return ErrCategoryNotFound
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  categoryID,
	}

	return u.repo.Create(context.Background(), product)
}

func (u *productUsecase) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	products, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *productUsecase) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidProductID
	}

	return u.repo.FindByID(ctx, objectID)
}

func (u *productUsecase) GetProductByName(ctx context.Context, name string) (*domain.Product, error) {
	if name == "" {
		return nil, ErrInvalidProductName
	}

	product, err := u.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

func (u *productUsecase) UpdateProduct(ctx context.Context, id string, req UpdateProductRequest) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidProductID
	}

	product, err := u.repo.FindByID(ctx, objectID)
	if err != nil {
		return err
	}
	if product == nil {
		return ErrProductNotFound
	}

	categoryID, err := primitive.ObjectIDFromHex(req.CategoryID)
	if err != nil {
		return ErrInvalidCategoryID
	}

	category, err := u.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return err
	}
	if category == nil {
		return ErrCategoryNotFound
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.CategoryID = categoryID

	err = u.repo.Update(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
