package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Items     []CartItem         `bson:"items" json:"items"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type CartItem struct {
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	Name      string             `bson:"name" json:"name"`
	Price     float64            `bson:"price" json:"price"`
	Quantity  int                `bson:"quantity" json:"quantity"`
}
