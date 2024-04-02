package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"samples-golang/model/request"
	"samples-golang/model/response"
	"samples-golang/repository"
	"samples-golang/utils"
)

type PostService struct {
	PostRepository repository.IPostRepository
}

type IPostService interface {
	HandleCreatedPost(e echo.Context, data request.RequestPost) (*response.CommonPostResponse, error)
	HandleGetAllPosts(e echo.Context) ([]*response.CommonPostResponse, error)
	HandleGetPostById(e echo.Context, postId string) (*response.CommonPostResponse, error)
	HandleRemovePostById(e echo.Context, postId string) error
	HandleUpdatePostById(e echo.Context, postId string, input request.UpdatePost) (*response.CommonPostResponse, error)
}

func NewPostService(postRepo repository.IPostRepository) IPostService {
	return PostService{
		PostRepository: postRepo,
	}
}

func (p PostService) HandleCreatedPost(ctx echo.Context, data request.RequestPost) (*response.CommonPostResponse, error) {
	tokenString := ctx.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		log.Errorf("Get payload failed: %s", err.Error())
		return nil, errors.New("Get payload failed. Please try again later")
	}

	userId, ok := payload["uid"].(string)
	if !ok {
		log.Errorf("User ID retrieval failed: %s", err.Error())
		return nil, errors.New("Sorry, couldn't retrieve user ID. Please try again")
	}

	dataResponse, err := p.PostRepository.CreatedPost(ctx.Request().Context(), data, userId)
	if err != nil {
		log.Errorf("Failed to create post: %s", err.Error())
		return nil, errors.New("Failed to create post. Please try again.")
	}

	return dataResponse, nil
}

func (p PostService) HandleGetAllPosts(ctx echo.Context) ([]*response.CommonPostResponse, error) {
	posts, err := p.PostRepository.GetAllPosts(ctx.Request().Context())
	if err != nil {
		log.Errorf("Failed to retrieve posts: %s", err.Error())
		return nil, errors.New("Failed to retrieve posts. Please try again later.")
	}

	return posts, nil
}

func (p PostService) HandleGetPostById(ctx echo.Context, postID string) (*response.CommonPostResponse, error) {
	post, err := p.PostRepository.GetPostById(ctx.Request().Context(), postID)
	if err != nil {
		log.Errorf("Failed to retrieve posts: %s", err.Error())
		return nil, errors.New("Failed to retrieve post by ID. Please try again later..")
	}

	return post, nil
}

func (p PostService) HandleRemovePostById(ctx echo.Context, postId string) error {
	tokenString := ctx.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		log.Errorf("Get payload failed: %s", err.Error())
		return errors.New("Get payload failed. Please try again later")
	}

	userId, ok := payload["uid"].(string)
	if !ok {
		log.Errorf("User ID retrieval failed: %s", err.Error())
		return errors.New("Sorry, couldn't retrieve user ID. Please try again")
	}

	err = p.PostRepository.RemovePostById(ctx.Request().Context(), postId, userId)
	if err != nil {
		log.Errorf("Failed to delete post: %s", err.Error())
		return errors.New("Failed to delete post. Please try again later.")
	}

	return nil
}

func (p PostService) HandleUpdatePostById(ctx echo.Context, postId string, input request.UpdatePost) (*response.CommonPostResponse, error) {
	tokenString := ctx.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		log.Errorf("Get payload failed: %s", err.Error())
		return nil, errors.New("Get payload failed. Please try again later")
	}

	userId, ok := payload["uid"].(string)
	if !ok {
		log.Errorf("User ID retrieval failed: %s", err.Error())
		return nil, errors.New("Sorry, couldn't retrieve user ID. Please try again")
	}

	updatedPost, err := p.PostRepository.UpdatePostById(ctx.Request().Context(), postId, input, userId)
	if err != nil {
		log.Errorf("Failed to update post: %s", err.Error())
		return nil, errors.New("Failed to update post. Please try again later.")
	}

	return updatedPost, nil
}
