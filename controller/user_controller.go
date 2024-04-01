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

func (a *UserController) GetAllUsers(c echo.Context) error {
	users, err := a.UserService.HandleGetAllUsers(c)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Data query successful",
		Data:       users,
	})
}

func (a *UserController) GetUserById(c echo.Context) error {
	userID := c.Param("id")

	user, err := a.UserService.HandleGetUserById(c, userID)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Search data successfully",
		Data:       user,
	})
}

func (a *UserController) UpdateUserById(c echo.Context) error {
	userId := c.Param("id")

	var input request.UpdateUser
	if err := c.Bind(&input); err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	if err := input.Validate(); err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	UserUpdated, err := a.UserService.HandleUpdateUserById(c, userId, input)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully updated data",
		Data:       UserUpdated,
	})
}

func (a *UserController) RemoveUserById(c echo.Context) error {
	userId := c.Param("userId")

	err := a.UserService.HandleRemoveUser(c, userId)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "user has been successfully removed",
		Data:       nil,
	})
}
