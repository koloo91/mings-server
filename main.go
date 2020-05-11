package main

import (
	"github.com/koloo91/mings-server/controller"
	"github.com/koloo91/mings-server/repository"
	"github.com/koloo91/mings-server/service"
	"log"
)

func main() {
	router := controller.SetupRoutes()

	//service.Repository = repository.NewInMemoryRepository()
	service.Repository = repository.NewFileRepository("storage")

	log.Fatal(router.Run(":9000"))
}
