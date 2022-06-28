package main

import (
	"fmt"

	"github.com/0l1v3rr/todo/app/controller"
	"github.com/0l1v3rr/todo/app/data"
	_ "github.com/0l1v3rr/todo/app/docs"
	"github.com/0l1v3rr/todo/app/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Advanced ToDo application
// @version         1.0
// @description     This is the API of the ToDo application

// @contact.name	API Support
// @contact.url 	https://0l1v3rr.github.io
// @contact.email 	oliver.mrakovics@gmail.com

// @license.name 	MIT
// @license.url 	https://opensource.org/licenses/MIT

// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// loading the environment variables
	err := data.GetVariables()
	if err != nil {
		fmt.Println("An error occurred while reading the .env file: ")
		fmt.Println(err.Error())
		return
	}

	// connecting to the db
	err = model.Setup()
	if err != nil {
		fmt.Println("Failed to connect to the database: ")
		fmt.Println(err.Error())
		return
	}

	// creating the gin router
	r := gin.Default()

	// using the cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// user endpoints
	r.GET("/api/v1/user", controller.GetLoggedInUser)
	r.POST("/api/v1/register", controller.Register)
	r.POST("/api/v1/login", controller.Login)
	r.POST("/api/v1/logout", controller.Logout)

	// task enpoints
	r.GET("/api/v1/tasks/list/:listId", controller.GetTasksByListId)
	r.GET("/api/v1/tasks/:id", controller.GetTaskById)
	r.POST("/api/v1/tasks", controller.CreateTask)
	r.PATCH("/api/v1/tasks/:id", controller.ChangeTaskStatus)
	r.PUT("/api/v1/tasks/:id", controller.EditTask)
	r.DELETE("/api/v1/tasks/:id", controller.DeleteTask)

	// swagger init
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// running the router
	r.Run(fmt.Sprintf(":%s", data.Env["PORT"]))
}
