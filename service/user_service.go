package service

import (
	"fmt"
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
	HandleGetUserById(c echo.Context, roomID string) (*model.User, error)
	HandleUpdateUserById(c echo.Context, userId string, input request.UpdateUser) (*model.User, error)
	HandleRemoveUser(c echo.Context, id string) error
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return UserService{
		UserRepository: userRepo,
	}
}

func (a UserService) HandleGetAllUsers(e echo.Context) ([]*model.User, error) {
	users, err := a.UserRepository.GetAllUsers(e.Request().Context())
	if err != nil {
		log.Errorf("User ID retrieval failed: %s", err.Error())
		return nil, errors.New("Sorry, couldn't retrieve user ID. Please try again")
	}

	return users, nil
}

func (a UserService) HandleGetUserById(c echo.Context, roomID string) (*model.User, error) {
	user, err := a.UserRepository.GetUserById(c.Request().Context(), roomID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) HandleUpdateUserById(c echo.Context, userId string, input request.UpdateUser) (*model.User, error) {
	roomUpdated, err := u.UserRepository.UpdateUserById(c.Request().Context(), userId, input)
	if err != nil {
		return nil, err
	}
	fmt.Println("service:", input.Email)

	return roomUpdated, nil
}

func (u UserService) HandleRemoveUser(c echo.Context, id string) error {
	err := u.UserRepository.RemoveRoomById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return nil
}
