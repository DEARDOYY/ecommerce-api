package repository

import (
	"context"
	"ecommerce-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
}

type categoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(collection *mongo.Collection) CategoryRepository {
	return &categoryRepository{
		collection: collection,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	result, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	category.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
