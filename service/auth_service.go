package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"samples-golang/repository"
	"samples-golang/utils"
)

type AuthService struct {
	AuthRepository repository.IAuthRepository
}

type IAuthService interface {
	HandleSignUp(ctx context.Context, tokenString string) (string, error)
	HandleLogin(c echo.Context) (string, error)
}

func NewAuthService(authRepo repository.IAuthRepository) IAuthService {
	return AuthService{
		AuthRepository: authRepo,
	}
}

func (u AuthService) HandleSignUp(ctx context.Context, tokenString string) (string, error) {
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		log.Errorf("Get payload failed: %s", err.Error())
		return "", errors.New("Get payload failed. Please try again later")
	}
	email, ok := payload["sub"].(string)
	if !ok {
		log.Errorf("Get email failed: %s", err.Error())
		return "", errors.New("Something happened. Please retry")
	}

	user, err := u.AuthRepository.InsertUser(ctx, email)
	if err != nil {
		log.Errorf("fail to insert user: %s", err.Error())
		return "", errors.New("Failed to insert user. Please try again later")
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		log.Errorf("Get token failed: %s", err.Error())
		return "", errors.New("Token creation failed. Please try again later")
	}

	return token, nil
}

func (u AuthService) HandleLogin(ctx echo.Context) (string, error) {
	tokenString := ctx.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		log.Errorf("Get payload failed: %s", err.Error())
		return "", errors.New("Get payload failed. Please try again later")
	}
	email, ok := payload["sub"].(string)
	if !ok {
		log.Errorf("Get email failed: %s", err.Error())
		return "", errors.New("get email failed. Please try again later")
	}
	user, err := u.AuthRepository.VerifyUser(ctx.Request().Context(), email)
	if err != nil {
		log.Errorf("User verification failed: %s", err.Error())
		return "", errors.New("User verification failed. Please double-check your information and try again.")
	}
	token, err := utils.GenerateToken(*user)
	if err != nil {
		log.Errorf("Get token failed: %s", err.Error())
		return "", errors.New("Token creation failed. Please try again later")
	}
	return token, nil
}
