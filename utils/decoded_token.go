package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"samples-golang/initializer"
	"samples-golang/model"
	"github.com/labstack/gommon/log"
)

func DecodeToken(tokenString string) (*model.JWTCustomsClaims, error) {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	if tokenString == "" {
		return nil, errors.New("invalid token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &model.JWTCustomsClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*model.JWTCustomsClaims)
	if !ok {
		return nil, err
	}

	if token.Claims.Valid() != nil {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
