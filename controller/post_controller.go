package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"samples-golang/model/request"
	"samples-golang/model/response"
	"samples-golang/service"
	"samples-golang/utils"
)

type PostController struct {
	PostService service.IPostService
}

func (p *PostController) CreatePost(ctx echo.Context) error {
	input := request.RequestPost{}
	if err := ctx.Bind(&input); err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	if err := input.Validate(); err != nil {
		utils.HandlerError(ctx, http.StatusForbidden, err.Error())
		return nil
	}

	data, err := p.PostService.HandleCreatedPost(ctx, input)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data query successful",
		Data:       data,
	})
}

func (p *PostController) GetAllPosts(ctx echo.Context) error {
	posts, err := p.PostService.HandleGetAllPosts(ctx)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
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
			ID:     post.ID,
			Title:  post.Title,
			Desc:   post.Description,
			UserId: post.UserId,
		})
	}
	
	type PostDataResponse struct {
		TotalPages int                    `json:"totalPages"`
		TotalItems int                    `json:"totalItems"`
		Posts      []*response.PostDataEntry `json:"posts"`
	}
	responseData := PostDataResponse{
		TotalPages: totalPages,
		TotalItems: totalItems,
		Posts:      docs,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

func (p *PostController) GetPostById(ctx echo.Context) error {
	postId := ctx.Param("id")

	data, err := p.PostService.HandleGetPostById(ctx, postId)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data query successful",
		Data:       data,
	})
}

func (p *PostController) RemovePostById(ctx echo.Context) error {
	postId := ctx.Param("id")

	err := p.PostService.HandleRemovePostById(ctx, postId)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "post has been successfully removed",
		Data:       nil,
	})

}

func (p *PostController) UpdatePostById(c echo.Context) error {
	postId := c.Param("postId")

	var input request.UpdatePost
	if err := c.Bind(&input); err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	if err := input.Validate(); err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	updatedPost, err := p.PostService.HandleUpdatePostById(c, postId, input)
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
