package models

import (
	"goapi/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User Struct
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty" binding:"required"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
}

// GetCollection function
func (u *User) GetCollection() *mongo.Collection {
	return db.Connection.Collection("users")
}
