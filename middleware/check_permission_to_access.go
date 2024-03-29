package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	verifier "github.com/okta/okta-jwt-verifier-golang"
	"net/http"
	"samples-golang/initializer"
	"samples-golang/model/response"
	"samples-golang/utils"
)

func CheckPermissionToAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				c.JSON(http.StatusForbidden, response.Response{
					StatusCode: http.StatusForbidden,
					Message:    "You are not signed in",
					Data:       nil,
				})
				return nil
			}
			err := verifyToken(tokenString)
			// TODO: Replace fmt with log (Zap)
			fmt.Println("err::::", err)
			if err != nil {
				c.JSON(http.StatusForbidden, response.Response{
					StatusCode: http.StatusForbidden,
					Message:    "Invalid token",
					Data:       nil,
				})
				return nil
			}
			return next(c)
		}
	}
}

func verifyToken(tokenString string) error {
	config, err := initializer.LoadConfig(".")

	tv := map[string]string{}
	tv["aud"] = "api://default"
	tv["cid"] = config.SpaClientId
	jv := verifier.JwtVerifier{
		Issuer:           config.Issuer,
		ClaimsToValidate: tv,
	}

	_, err = jv.New().VerifyAccessToken(tokenString)
	if err != nil {
		return err
	}

	return nil
}

func CheckPermissionToAccessByRole(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				c.JSON(http.StatusForbidden, response.Response{
					StatusCode: http.StatusForbidden,
					Message:    "You are not signed in",
					Data:       nil,
				})
				return nil
			}
			claims, _ := utils.DecodeToken(tokenString)
			if claims.Role != requiredRole {
				c.JSON(http.StatusForbidden, response.Response{
					StatusCode: http.StatusForbidden,
					Message:    "You don't have permission to access on this feature",
					Data:       nil,
				})
				return nil
			}
			return next(c)
		}
	}
}