package repository

import (
	"golang.org/x/net/context"
	"samples-golang/model"
)

type IAuthRepository interface {
	InsertUser(context context.Context, email string) (*model.User, error)
	VerifyUser(context context.Context, email string) (*model.User, error)
}
