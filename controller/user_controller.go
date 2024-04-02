package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"samples-golang/model/request"
	"samples-golang/model/response"
	"samples-golang/service"
	"samples-golang/utils"
)

type UserController struct {
	UserService service.IUserService
}

func (u *UserController) GetAllUsers(ctx echo.Context) error {
	users, err := u.UserService.HandleGetAllUsers(ctx)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data query successful",
		Data:       users,
	})
}

func (u *UserController) GetUserById(ctx echo.Context) error {
	userID := ctx.Param("id")

	user, err := u.UserService.HandleGetUserById(ctx, userID)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Search data successfully",
		Data:       user,
	})
}

func (u *UserController) UpdateUserById(ctx echo.Context) error {
	userId := ctx.Param("id")

	var input request.UpdateUser
	if err := ctx.Bind(&input); err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	if err := input.Validate(); err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	UserUpdated, err := u.UserService.HandleUpdateUserById(ctx, userId, input)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully updated data",
		Data:       UserUpdated,
	})
}

func (u *UserController) RemoveUserById(ctx echo.Context) error {
	userId := ctx.Param("userId")

	err := u.UserService.HandleRemoveUser(ctx, userId)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "user has been successfully removed",
		Data:       nil,
	})
}
