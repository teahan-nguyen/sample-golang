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
	HandleCreatedPost(e echo.Context, data request.ReqPost) (*response.CommonPostResponse, error)
	HandleGetAllPosts(e echo.Context) ([]*response.CommonPostResponse, error)
	HandleGetPostById(e echo.Context, postId string) (*response.CommonPostResponse, error)
	HandleRemovePostById(e echo.Context, postId string) error
	HandleUpdatePostById(e echo.Context, postId string, input request.ReqUpdatePost) (*response.CommonPostResponse, error)
}

func NewPostService(postRepo repository.IPostRepository) IPostService {
	return PostService{
		PostRepository: postRepo,
	}
}

func (a PostService) HandleCreatedPost(e echo.Context, data request.ReqPost) (*response.CommonPostResponse, error) {
	tokenString := e.Request().Header.Get("Authorization")
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

	dataResponse, err := a.PostRepository.CreatedPost(e.Request().Context(), data, userId)
	if err != nil {
		log.Errorf("Failed to create post: %s", err.Error())
		return nil, errors.New("Failed to create post. Please try again.")
	}

	return dataResponse, nil
}

func (a PostService) HandleGetAllPosts(e echo.Context) ([]*response.CommonPostResponse, error) {
	posts, err := a.PostRepository.GetAllPosts(e.Request().Context())
	if err != nil {
		log.Errorf("Failed to retrieve posts: %s", err.Error())
		return nil, errors.New("Failed to retrieve posts. Please try again later.")
	}

	return posts, nil
}

func (a PostService) HandleGetPostById(e echo.Context, postID string) (*response.CommonPostResponse, error) {
	post, err := a.PostRepository.GetPostById(e.Request().Context(), postID)
	if err != nil {
		log.Errorf("Failed to retrieve posts: %s", err.Error())
		return nil, errors.New("Failed to retrieve post by ID. Please try again later..")
	}

	return post, nil
}

func (a PostService) HandleRemovePostById(e echo.Context, postId string) error {
	tokenString := e.Request().Header.Get("Authorization")
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

	err = a.PostRepository.RemovePostById(e.Request().Context(), postId, userId)
	if err != nil {
		log.Errorf("Failed to delete post: %s", err.Error())
		return errors.New("Failed to delete post. Please try again later.")
	}

	return nil
}

func (a PostService) HandleUpdatePostById(e echo.Context, postId string, input request.ReqUpdatePost) (*response.CommonPostResponse, error) {
	tokenString := e.Request().Header.Get("Authorization")
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

	updatedPost, err := a.PostRepository.UpdatePostById(e.Request().Context(), postId, input, userId)
	if err != nil {
		log.Errorf("Failed to update post: %s", err.Error())
		return nil, errors.New("Failed to update post. Please try again later.")
	}

	return updatedPost, nil
}
