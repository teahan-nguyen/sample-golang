package repository

import (
	"golang.org/x/net/context"
	"samples-golang/model"
	"samples-golang/model/request"
	"samples-golang/model/response"
)

type AutheRepository interface {
	CreatedPost(context context.Context, data request.ReqPost, userId string) (*response.ResPostData, error)
	GetAllPosts(context context.Context) ([]*response.ResPostData, error)
	GetPostById(context context.Context, postId string) (*response.ResPostData, error)
	RemovePostById(context context.Context, postId string, userId string) error
	UpdatePostById(context context.Context, postId string, input response.ResPostData, userId string) (*response.ResPostData, error)
	InsertUser(context context.Context, email string) (*model.User, error)
	VerifyUser(context context.Context, email string) (*model.User, error)
}
