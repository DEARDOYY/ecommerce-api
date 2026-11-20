package repository

import (
	"context"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ecommerce-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"errors"
)

type ProductRepository interface {
	EnsureIndexes(ctx context.Context) error
	Create(ctx context.Context, product *domain.Product) error
	FindAll(ctx context.Context) ([]*domain.Product, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error)
	FindByName(ctx context.Context, name string) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	Search(ctx context.Context, f domain.ProductFilter) ([]*domain.Product, int64, error)
}

var ErrProductExists = errors.New("product already exists")

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
	if err != nil {
		return err
	}

	_, err = r.collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "category_id", Value: 1}, {Key: "price", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: -1}}},
	})
	if err != nil {
		return err
	}
	return nil
}

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

func (r *productRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
	var product domain.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) FindByName(ctx context.Context, name string) (*domain.Product, error) {
	var product domain.Product
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"is_active": status}})
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Search(ctx context.Context, f domain.ProductFilter) ([]*domain.Product, int64, error) {
	filter := bson.M{"is_active": true}

	if f.Q != "" {
		pattern := regexp.QuoteMeta(f.Q)
		filter["$or"] = bson.A{
			bson.M{"name": bson.M{"$regex": pattern, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": pattern, "$options": "i"}},
		}
	}

	if f.CategoryID != nil {
		filter["category_id"] = f.CategoryID
	}

	price := bson.M{}
	if f.MinPrice != nil {
		price["$gte"] = *f.MinPrice
	}
	if f.MaxPrice != nil {
		price["$lte"] = *f.MaxPrice
	}
	if len(price) > 0 {
		filter["price"] = price
	}

	if f.InStock {
		filter["stock"] = bson.M{"$gt": 0}
	}

	// sort
	sort := bson.D{{Key: "created_at", Value: -1}}
	switch f.Sort {
	case "price_asc":
		sort = bson.D{{Key: "price", Value: 1}}
	case "price_desc":
		sort = bson.D{{Key: "price", Value: -1}}
	}

	opts := options.Find().
		SetSort(sort).
		SetSkip((f.Page - 1) * f.Limit).
		SetLimit(f.Limit)

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
