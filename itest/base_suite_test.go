package itest

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koloo91/mings-server/controller"
	"github.com/koloo91/mings-server/repository"
	"github.com/koloo91/mings-server/service"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type ResourceFile = string

const (
	usrFile ResourceFile = "usr.yml"
)

type BaseSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *BaseSuite) SetupSuite() {
	log.Println("Setup suite")
	suite.router = controller.SetupRoutes()
}

func (suite *BaseSuite) SetupTest() {
	log.Println("Setup test")

	os.RemoveAll("./storage")

	//service.Repository = repository.NewInMemoryRepository()
	service.Repository = repository.NewFileRepository("./storage")
}

func (suite *BaseSuite) TearDownTest() {
	log.Println("Tear down test")
	os.RemoveAll("./storage")
}

func TestBaseSuiteTest(t *testing.T) {
	suite.Run(t, &BaseSuite{})
}

func prepareMultipartUpload(resourceFile ResourceFile) (*bytes.Buffer, *multipart.Writer) {
	file, err := os.Open(fmt.Sprintf("./resources/%s", resourceFile))
	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "file")
	if err != nil {
		log.Fatal(err)
	}
	part.Write(fileContent)

	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	return body, writer
}

func (suite *BaseSuite) uploadDocument(resourceFile ResourceFile) {
	multipartBody, writer := prepareMultipartUpload(resourceFile)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/documents", multipartBody)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusOK, recorder.Code)
}
