package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"samples-golang/controller"
	"samples-golang/db"
	"samples-golang/initializer"
	"samples-golang/repository/repo_implement"
	"samples-golang/routes"
	"samples-golang/service"
)

func main() {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	mongoDB := &db.MongoDB{
		DbName: config.DbName,
	}

	mongoDB.Connect()
	defer mongoDB.Close()

	e := echo.New()
	e.Use()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.ClientOrigin2, config.ClientOrigin},
		AllowCredentials: true,
	}))

	postRepo := repo_implement.NewPostImplement(mongoDB.Client.Database(config.DbName))
	postService := service.NewPostService(postRepo)
	postController := controller.PostController{
		PostService: postService,
	}

	authRepo := repo_implement.NewAuthRepository(mongoDB.Client.Database(config.DbName))
	authService := service.NewAuthService(authRepo)
	authController := controller.AuthController{
		AuthService: authService,
	}

	userRepo := repo_implement.NewUserImplement(mongoDB.Client.Database(config.DbName))
	userService := service.NewUserService(userRepo)
	userController := controller.UserController{
		UserService: userService,
	}

	e.GET("/healthchecker", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "success",
			"message": "Welcome to Golang",
		})
	})

	api := routes.API{
		Echo:           e,
		PostController: postController,
		AuthController: authController,
		UserController: userController,
	}
	api.SetUpRouter()

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
