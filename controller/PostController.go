package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"samples-golang/model/request"
	"samples-golang/model/response"
	"samples-golang/service"
	"samples-golang/utils"
)

type PostController struct {
	PostService service.PostService
}

func (a *PostController) CreatedPost(c echo.Context) error {
	input := request.ReqPost{}
	if err := c.Bind(&input); err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		utils.HandlerError(c, http.StatusForbidden, err.Error())
		return nil
	}

	data, err := a.PostService.HandleCreatedPost(c, input)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data query successful",
		Data:       data,
	})
}

func (a *PostController) GetAllPosts(c echo.Context) error {
	posts, err := a.PostService.HandleGetAllPosts(c)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	pageSize := 10
	totalItems := len(posts)
	totalPages := totalItems / pageSize
	if totalItems%pageSize != 0 {
		totalPages++
	}

	var docs []*response.PostDataEntry
	for _, post := range posts {
		docs = append(docs, &response.PostDataEntry{
			ID:    post.ID,
			Title: post.Title,
			Desc:  post.Description,
		})
	}
	responseData := struct {
		TotalPages int                       `json:"totalPages"`
		TotalItems int                       `json:"totalItems"`
		Docs       []*response.PostDataEntry `json:"docs"`
	}{
		TotalPages: totalPages,
		TotalItems: totalItems,
		Docs:       docs,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (a *PostController) GetPostById(c echo.Context) error {
	PostId := c.Param("id")

	data, err := a.PostService.HandleGetPostById(c, PostId)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data query successful",
		Data:       data,
	})
}

func (a *PostController) RemovePostById(c echo.Context) error {
	id := c.Param("id")

	err := a.PostService.HandleRemovePostById(c, id)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "post has been successfully removed",
		Data:       nil,
	})

}

func (a *PostController) UpdatePostById(c echo.Context) error {
	id := c.Param("id")

	var input response.ResPostData
	if err := c.Bind(&input); err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	updatedPost, err := a.PostService.HandleUpdatePostById(c, id, input)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully updated data",
		Data:       updatedPost,
	})
}
