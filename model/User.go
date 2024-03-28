package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email"`
	Role  string             `json:"role" bson:"role"`
}
