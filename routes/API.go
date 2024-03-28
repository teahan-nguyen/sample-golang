package routes

import (
	"github.com/labstack/echo/v4"
	"samples-golang/controller"
	"samples-golang/middleware"
)

type API struct {
	Echo           *echo.Echo
	PostController controller.PostController
	AuthController controller.AuthController
}

func (API *API) SetUpRouter() {
	//Post
	post := API.Echo.Group("/api/v1/posts", middleware.CheckPermissionToAccess())
	post.POST("", API.PostController.CreatedPost)
	post.GET("", API.PostController.GetAllPosts)
	post.GET("/:id", API.PostController.GetPostById)
	post.DELETE("/:id", API.PostController.RemovePostById)
	post.PUT("/:id", API.PostController.UpdatePostById)

	auth := API.Echo.Group("/api/v1", middleware.CheckPermissionToAccess())
	auth.POST("/login", API.AuthController.Login)

	API.Echo.POST("/signUp", API.AuthController.SignUp)
}
