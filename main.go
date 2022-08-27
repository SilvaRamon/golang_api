package main

import (
	"hello/controllers"
	"hello/repository"
	"hello/service"
	"hello/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage := storage.Initialize()

	router := gin.Default()

	repository := repository.Initialize(*storage)

	service := service.Initialize(repository)

	controllers.Initialize(service, router).ExposeRoutes()

	router.Run(":8081")
}
