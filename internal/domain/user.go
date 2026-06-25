package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	Role      string             `bson:"role" json:"role"` // "user" | "admin"
	Address   []Address          `bson:"address" json:"address"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Address struct {
	Label     string `bson:"label" json:"label"` // "บ้าน", "ที่ทำงาน"
	Address   string `bson:"address" json:"address"`
	Phone     string `bson:"phone" json:"phone"`
	IsDefault bool   `bson:"is_default" json:"is_default"`
}
