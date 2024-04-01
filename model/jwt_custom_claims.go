package model

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTCustomsClaims struct {
	ID    primitive.ObjectID
	Role  string
	Email string
	jwt.StandardClaims
}
