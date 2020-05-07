package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koloo91/mings-server/controller"
	"github.com/koloo91/mings-server/repository"
	"github.com/koloo91/mings-server/service"
	"log"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	setupRoutes(router)

	service.Repository = repository.NewInMemoryRepository()

	log.Fatal(router.Run(":9000"))
}

func setupRoutes(router *gin.Engine) {
	router.GET("/documents", controller.All)
	router.POST("/documents", controller.Upload)
	router.GET("/documents/:id", controller.ById)
}
