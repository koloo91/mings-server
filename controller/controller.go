package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koloo91/mings-server/model"
	"github.com/koloo91/mings-server/service"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
)

func All(ctx *gin.Context) {
	documents, err := service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"documents": documents})
}

func Upload(ctx *gin.Context) {
	multipartFile, err := ctx.FormFile("file")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error while uploading file"})
		return
	}
	file, err := multipartFile.Open()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error while uploading file"})
		return
	}
	defer file.Close()

	var document model.Document
	if err := yaml.NewDecoder(file).Decode(&document); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse yaml file"})
		return
	}

	err = service.Create(document)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable store document"})
		return
	}

	ctx.JSON(http.StatusOK, document)
}

func ById(ctx *gin.Context) {

}
