package utils

import (
	"github.com/labstack/echo/v4"
	"samples-golang/model/response"
)

func HandlerError(c echo.Context, StatusCode int, Message string) {
	c.JSON(StatusCode, response.Response{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       nil,
	})
}
