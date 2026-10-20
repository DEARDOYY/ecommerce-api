package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ecommerce-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"errors"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	EnsureIndexes(ctx context.Context) error
}

type categoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(collection *mongo.Collection) CategoryRepository {
	return &categoryRepository{
		collection: collection,
	}
}

func (r *categoryRepository) EnsureIndexes(ctx context.Context) error {
	_, err := r.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "slug", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

var ErrCategoryExists = errors.New("category already exists")

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	result, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return ErrCategoryExists
		}
		return err
	}

	category.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
