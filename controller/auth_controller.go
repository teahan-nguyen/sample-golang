package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"samples-golang/model/response"
	"samples-golang/service"
	"samples-golang/utils"
)

type AuthController struct {
	AuthService service.IAuthService
}

func (a *AuthController) SignUp(ctx echo.Context) error {
	tokenString := ctx.Request().Header.Get("Authorization")
	newUser, err := a.AuthService.HandleSignUp(ctx.Request().Context(), tokenString)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Account created successfully",
		Data:       newUser,
	})
}

func (a *AuthController) Login(ctx echo.Context) error {
	token, err := a.AuthService.HandleLogin(ctx)
	if err != nil {
		utils.HandlerError(ctx, http.StatusBadRequest, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Login successful",
		Data:       token,
	})
}
