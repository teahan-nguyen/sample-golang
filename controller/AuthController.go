package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"samples-golang/model/response"
	"samples-golang/service"
	"samples-golang/utils"
)

type AuthController struct {
	AuthService service.AuthService
}

func (a *AuthController) SignUp(c echo.Context) error {
	newUser, err := a.AuthService.HandleSignUp(c)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Account created successfully",
		Data:       newUser,
	})
}

func (a *AuthController) Login(c echo.Context) error {
	token, err := a.AuthService.HandleLogin(c)
	if err != nil {
		utils.HandlerError(c, http.StatusBadRequest, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Login successful",
		Data:       token,
	})
}
