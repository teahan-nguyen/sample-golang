package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"samples-golang/controller"
	"samples-golang/initializer"
	"samples-golang/middleware"
)

type API struct {
	Echo           *echo.Echo
	PostController controller.PostController
	AuthController controller.AuthController
	UserController controller.UserController
}

func (API *API) SetUpRouter() {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	//Post
	apiV1 := API.Echo.Group("api/v1")
	post := apiV1.Group("/posts", middleware.CheckPermissionToAccess())
	post.POST("", API.PostController.CreatePost)
	post.GET("", API.PostController.GetAllPosts)
	post.GET("/:id", API.PostController.GetPostById)
	post.DELETE("/:id", API.PostController.RemovePostById)
	post.PUT("/:id", API.PostController.UpdatePostById)
	//User
	user := apiV1.Group("/users", middleware.CheckPermissionToAccessByRole(config.Role))
	user.GET("", API.UserController.GetAllUsers)
	user.GET("/:id", API.UserController.GetUserById)
	user.DELETE("/:id", API.UserController.RemoveUserById)
	user.PUT("/:id", API.UserController.UpdateUserById)

	auth := apiV1.Group("/api/v1", middleware.CheckPermissionToAccess())
	auth.POST("/login", API.AuthController.Login)

	API.Echo.POST("/signUp", API.AuthController.SignUp)

}
