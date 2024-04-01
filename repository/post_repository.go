package repository

import (
	"golang.org/x/net/context"
	"samples-golang/model/request"
	"samples-golang/model/response"
)

type IPostRepository interface {
	CreatedPost(context context.Context, data request.ReqPost, userId string) (*response.CommonPostResponse, error)
	GetAllPosts(context context.Context) ([]*response.CommonPostResponse, error)
	GetPostById(context context.Context, postId string) (*response.CommonPostResponse, error)
	RemovePostById(context context.Context, postId string, userId string) error
	UpdatePostById(context context.Context, postId string, input request.ReqUpdatePost, userId string) (*response.CommonPostResponse, error)
}
