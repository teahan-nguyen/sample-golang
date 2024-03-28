package service

import (
	"github.com/labstack/echo/v4"
	"samples-golang/repository"
	"samples-golang/utils"
)

type AuthService struct {
	AuthRepository repository.AutheRepository
}

func (u *AuthService) HandleSignUp(c echo.Context) (string, error) {
	tokenString := c.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		return "", err
	}
	email, ok := payload["sub"].(string)
	if !ok {
		return "", err
	}

	user, err := u.AuthRepository.InsertUser(c.Request().Context(), email)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthService) HandleLogin(c echo.Context) (string, error) {
	tokenString := c.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		return "", err
	}
	email, ok := payload["sub"].(string)
	if !ok {
		return "", err
	}
	user, err := u.AuthRepository.VerifyUser(c.Request().Context(), email)
	token, err := utils.GenerateToken(*user)
	return token, nil
}
