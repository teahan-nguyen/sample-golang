package repository

import (
	"golang.org/x/net/context"
	"samples-golang/model"
	"samples-golang/model/request"
)

type IUserRepository interface {
	GetAllUsers(context context.Context) ([]*model.User, error)
	GetUserById(context context.Context, userId string) (*model.User, error)
	UpdateUserById(context context.Context, id string, input request.UpdateUser) (*model.User, error)
	RemoveUserById(context context.Context, userId string) error
}
