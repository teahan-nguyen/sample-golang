package service

import (
	"github.com/labstack/echo/v4"
	"samples-golang/model/request"
	"samples-golang/model/response"
	"samples-golang/repository"
	"samples-golang/utils"
)

type PostService struct {
	PostRepository repository.AutheRepository
}

func (a *PostService) HandleCreatedPost(e echo.Context, data request.ReqPost) (*response.ResPostData, error) {
	tokenString := e.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		return nil, err
	}
	userId, ok := payload["uid"].(string)
	if !ok {
		return nil, err
	}

	dataResponse, err := a.PostRepository.CreatedPost(e.Request().Context(), data, userId)
	if err != nil {
		return nil, err
	}

	return dataResponse, nil
}

func (a *PostService) HandleGetAllPosts(e echo.Context) ([]*response.ResPostData, error) {
	posts, err := a.PostRepository.GetAllPosts(e.Request().Context())
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (a *PostService) HandleGetPostById(e echo.Context, postID string) (*response.ResPostData, error) {
	post, err := a.PostRepository.GetPostById(e.Request().Context(), postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (a *PostService) HandleRemovePostById(e echo.Context, id string) error {
	tokenString := e.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		return err
	}
	userId, ok := payload["uid"].(string)
	if !ok {
		return err
	}

	err = a.PostRepository.RemovePostById(e.Request().Context(), id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (a *PostService) HandleUpdatePostById(e echo.Context, postId string, input response.ResPostData) (*response.ResPostData, error) {
	tokenString := e.Request().Header.Get("Authorization")
	payload, err := utils.GetTokenPayload(tokenString)
	if err != nil {
		return nil, err
	}
	userId, ok := payload["uid"].(string)
	if !ok {
		return nil, err
	}

	updatedPost, err := a.PostRepository.UpdatePostById(e.Request().Context(), postId, input, userId)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}
