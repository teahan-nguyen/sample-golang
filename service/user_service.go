package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"samples-golang/model"
	"samples-golang/model/request"
	"samples-golang/repository"
)

type UserService struct {
	UserRepository repository.IUserRepository
}
type IUserService interface {
	HandleGetAllUsers(e echo.Context) ([]*model.User, error)
	HandleGetUserById(c echo.Context, userId string) (*model.User, error)
	HandleUpdateUserById(c echo.Context, userId string, input request.UpdateUser) (*model.User, error)
	HandleRemoveUser(c echo.Context, userId string) error
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return UserService{
		UserRepository: userRepo,
	}
}

func (u UserService) HandleGetAllUsers(ctx echo.Context) ([]*model.User, error) {
	users, err := u.UserRepository.GetAllUsers(ctx.Request().Context())
	if err != nil {
		log.Errorf("User ID retrieval failed: %s", err.Error())
		return nil, errors.New("Sorry, couldn't retrieve user ID. Please try again")
	}

	return users, nil
}

func (u UserService) HandleGetUserById(ctx echo.Context, userId string) (*model.User, error) {
	user, err := u.UserRepository.GetUserById(ctx.Request().Context(), userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) HandleUpdateUserById(ctx echo.Context, userId string, input request.UpdateUser) (*model.User, error) {
	user, err := u.UserRepository.UpdateUserById(ctx.Request().Context(), userId, input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) HandleRemoveUser(ctx echo.Context, userId string) error {
	err := u.UserRepository.RemoveUserById(ctx.Request().Context(), userId)
	if err != nil {
		return err
	}
	return nil
}
