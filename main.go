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
		DbName: "testMongoDb",
	}

	mongoDB.Connect()
	defer mongoDB.Close()

	e := echo.New()
	e.Use()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.ClientOrigin2, config.ClientOrigin},
		AllowCredentials: true,
	}))

	postService := service.PostService{
		PostRepository: repo_implement.NewImplement(mongoDB),
	}
	postController := controller.PostController{
		PostService: postService,
	}

	authService := service.AuthService{
		AuthRepository: repo_implement.NewImplement(mongoDB),
	}
	authController := controller.AuthController{
		AuthService: authService,
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
	}
	api.SetUpRouter()

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
