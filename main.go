package main

import (
	"github.com/koloo91/mings-server/controller"
	"github.com/koloo91/mings-server/repository"
	"github.com/koloo91/mings-server/service"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	applicationLogWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "./logs/application.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
	})

	logrus.SetOutput(applicationLogWriter)
}

func main() {
	router := controller.SetupRoutes()

	server := &http.Server{
		Addr:         ":9000",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	//service.Repository = repository.NewInMemoryRepository()
	service.Repository = repository.NewFileRepository("./storage")

	logrus.Info("Starting server")
	logrus.Error(server.ListenAndServe())
}
