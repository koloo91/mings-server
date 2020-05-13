package repository

import (
	"encoding/json"
	"fmt"
	"github.com/koloo91/mings-server/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type FileRepository struct {
	baseDir      string
	refreshCache bool
	documents    map[string]model.Document
	mutex        sync.Mutex
}

func NewFileRepository(baseDir string) *FileRepository {
	return &FileRepository{
		baseDir:      baseDir,
		refreshCache: true,
		documents:    make(map[string]model.Document),
	}
}

func (r *FileRepository) Get() ([]model.Document, error) {
	returnValue := make([]model.Document, 0, len(r.documents))

	if _, err := os.Stat(fmt.Sprintf(r.baseDir)); os.IsNotExist(err) {
		return returnValue, nil
	}

	if r.refreshCache {
		r.loadData()
	}

	for _, document := range r.documents {
		returnValue = append(returnValue, document)
	}

	return returnValue, nil
}

func (r *FileRepository) loadData() {
	logrus.Info("Refreshing cache")
	fileInfos, err := ioutil.ReadDir(r.baseDir)
	if err != nil {
		logrus.Error(err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		fileName := fmt.Sprintf("%s/%s", r.baseDir, fileInfo.Name())
		file, err := os.Open(fileName)
		if err != nil {
			logrus.Error(err)
			continue
		}

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			logrus.Debug(err)
			continue
		}

		var document model.Document
		if err := json.Unmarshal(fileContent, &document); err != nil {
			logrus.Debug(err)
			continue
		}

		r.documents[document.Id] = document
	}

	r.refreshCache = false
}

func (r *FileRepository) Save(document model.Document) error {
	if _, err := os.Stat(fmt.Sprintf("%s", r.baseDir)); os.IsNotExist(err) {
		if err := os.Mkdir(r.baseDir, os.ModePerm); err != nil {
			return err
		}
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	documentBytes, err := json.Marshal(document)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.json", r.baseDir, document.Id), documentBytes, os.ModePerm); err != nil {
		return err
	}

	r.documents[document.Id] = document

	return nil
}

func (r *FileRepository) ById(id string) (model.Document, error) {
	if r.refreshCache {
		r.loadData()
	}

	document, exists := r.documents[id]

	if !exists {
		return document, ErrDocumentNotFound
	}

	return document, nil
}

func (r *FileRepository) GetByDependsOn(dependsOn string) ([]model.Document, error) {
	if r.refreshCache {
		r.loadData()
	}

	returnValue := make([]model.Document, 0, len(r.documents))
	dependsOnToLower := strings.ToLower(dependsOn)

	for _, document := range r.documents {
		if documentDependsOn(dependsOnToLower, document) {
			returnValue = append(returnValue, document)
		}
	}

	return returnValue, nil
}

func documentDependsOn(dependsOn string, document model.Document) bool {
	for _, dependency := range document.Service.DependsOn.Internal {
		if strings.ToLower(dependency.ServiceName) == dependsOn {
			return true
		}
	}

	for _, dependency := range document.Service.DependsOn.External {
		if strings.ToLower(dependency.ServiceName) == dependsOn {
			return true
		}
	}

	return false
}
