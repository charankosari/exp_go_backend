package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the schema for a user in MongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Age       int                `bson:"age,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Phone     string             `bson:"phone,omitempty"`
}
