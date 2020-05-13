package controller

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/koloo91/mings-server/docs"
	"github.com/koloo91/mings-server/model"
	"github.com/koloo91/mings-server/repository"
	"github.com/koloo91/mings-server/service"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"time"
)

// @title mings-server API
// @version 1.0
func SetupRoutes() *gin.Engine {
	logrus.Info("Setting up routes")

	router := gin.New()

	router.Use(Logger(), cors.Default(), gin.Recovery())

	router.GET("/documents", all)
	router.POST("/documents", upload)
	router.GET("/documents/:id", byId)

	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "GIN_MODE"))

	return router
}

func Logger() gin.HandlerFunc {
	accessLog := logrus.New()
	accessLog.SetFormatter(&logrus.JSONFormatter{})
	accessLogWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "./logs/access.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
	})

	accessLog.SetOutput(accessLogWriter)
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next()

		latency := time.Since(t)

		accessLog.WithFields(logrus.Fields{
			"code":        ctx.Writer.Status(),
			"duration_ns": latency,
			"method":      ctx.Request.Method,
			"ip":          ctx.ClientIP(),
			"path":        ctx.Request.URL.Path,
		}).Info("")
	}
}

// GetDocuments godoc
// @Summary Get all stored documents
// @Description Get all stored documents
// @ID get-documents
// @Produce json
// @Param depends_on query string false "depends on"
// @Success 200 {object} model.Documents
// @Failure 400 {object} model.ApiError
// @Failure 500 {object} model.ApiError
// @Router /documents [get]
func all(ctx *gin.Context) {
	if dependsOn, exists := ctx.GetQuery("depends_on"); exists {
		documents, err := service.GetByDependsOn(dependsOn)
		if err != nil {
			logrus.Error(err)
			ctx.JSON(http.StatusBadRequest, model.ApiError{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, model.Documents{Documents: documents})
	} else {
		documents, err := service.GetAll()
		if err != nil {
			logrus.Error(err)
			ctx.JSON(http.StatusBadRequest, model.ApiError{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, model.Documents{Documents: documents})
	}
}

// UploadDocument godoc
// @Summary Uploads a single document
// @Description Uploads a single document
// @ID upload-document
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "document file"
// @Success 200 {object} model.Document
// @Failure 400 {object} model.ApiError
// @Failure 500 {object} model.ApiError
// @Router /documents [post]
func upload(ctx *gin.Context) {
	multipartFile, err := ctx.FormFile("file")
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, model.ApiError{Message: "error while uploading file"})
		return
	}

	file, err := multipartFile.Open()
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, model.ApiError{Message: "error while uploading file"})
		return
	}
	defer file.Close()

	var document model.Document
	if err := yaml.NewDecoder(file).Decode(&document); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, model.ApiError{Message: "unable to parse yaml file"})
		return
	}

	err = service.Create(document)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, model.ApiError{Message: "unable store document"})
		return
	}

	ctx.JSON(http.StatusOK, document)
}

// GetDocumentById godoc
// @Summary Get all stored documents
// @Description Get all stored documents
// @ID get-document-by-id
// @Produce json
// @Param id path string true "document id"
// @Success 200 {object} model.Document
// @Failure 400 {object} model.ApiError
// @Failure 500 {object} model.ApiError
// @Router /documents/{id} [get]
func byId(ctx *gin.Context) {
	documentId := ctx.Param("id")

	document, err := service.GetById(documentId)
	if err != nil {
		logrus.Error(err)
		if err == repository.ErrDocumentNotFound {
			ctx.JSON(http.StatusNotFound, model.ApiError{Message: fmt.Sprintf("document with id '%s' not found", documentId)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, "unexpected error")
		return
	}

	ctx.JSON(http.StatusOK, document)
}
