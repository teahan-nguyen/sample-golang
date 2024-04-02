package utils

import (
	"github.com/golang-jwt/jwt"
	"samples-golang/initializer"
	"samples-golang/model"
	"github.com/labstack/gommon/log"
	"time"
)

func GenerateToken(user model.User) (string, error) {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	claims := &model.JWTCustomsClaims{
		ID:    user.Id,
		Role:  user.Role,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	results, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}

	return results, nil
}
