package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ecommerce-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"errors"
)

type ProductRepository interface{}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) ProductRepository {
	return &productRepository{
		collection: collection,
	}
}

func (r *productRepository) EnsureIndexes(ctx context.Context) error {
	_, err := r.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

var ErrProductExists = errors.New("product already exists")

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return ErrProductExists
		}
		return err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
