package routes

import (
	"github.com/labstack/echo/v4"
	"samples-golang/controller"
	"samples-golang/initializer"
	"samples-golang/middleware"
)

type API struct {
	Echo           *echo.Echo
	PostController controller.PostController
	AuthController controller.AuthController
	UserController controller.UserController

	//this
}

func (API *API) SetUpRouter() {
	config, _ := initializer.LoadConfig(".")
	//Post
	post := API.Echo.Group("/api/v1/posts", middleware.CheckPermissionToAccess())
	post.POST("", API.PostController.CreatePost)
	post.GET("", API.PostController.GetAllPosts)
	post.GET("/:id", API.PostController.GetPostById)
	post.DELETE("/:id", API.PostController.RemovePostById)
	post.PUT("/:id", API.PostController.UpdatePostById)
	//user
	user := API.Echo.Group("/api/v1/users", middleware.CheckPermissionToAccessByRole(config.Role))
	user.GET("", API.UserController.GetAllUsers)
	user.GET("/:id", API.UserController.GetUserById)
	user.DELETE("/:id", API.UserController.RemoveUserById)
	user.PUT("/:id", API.UserController.UpdateUserById)

	auth := API.Echo.Group("/api/v1", middleware.CheckPermissionToAccess())
	auth.POST("/login", API.AuthController.Login)

	API.Echo.POST("/signUp", API.AuthController.SignUp)

}
