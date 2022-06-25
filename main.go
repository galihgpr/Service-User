package main

import (
	"alta-test/config"
	"alta-test/controller/handler"
	"alta-test/controller/router"
	"alta-test/model"
	"alta-test/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	appConfig := config.Get()
	postgreDB := config.InitDB(appConfig)
	fmt.Println(postgreDB)
	config.MigrateDB(postgreDB)

	// Inject Model
	modelUser := model.NewModelDB(postgreDB)

	// Inject Service
	serviceUser := service.NewServiceModel(modelUser)

	// Inject Handler
	handlerUser := handler.NewUserHandler(serviceUser, validator.New())

	// Init GIN
	e := gin.New()

	// Use Router
	router.Router(e, handlerUser)

	// Run Apps
	e.Run(":8000")
}
