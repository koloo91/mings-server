package repository

import (
	"encoding/json"
	"fmt"
	"github.com/koloo91/mings-server/model"
	"io/ioutil"
	"log"
	"os"
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

	if _, err := os.Stat(fmt.Sprintf("./%s", r.baseDir)); os.IsNotExist(err) {
		return returnValue, nil
	}

	if r.refreshCache {
		fileInfos, err := ioutil.ReadDir(fmt.Sprintf("./%s", r.baseDir))
		if err != nil {
			return returnValue, err
		}

		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				continue
			}

			fileName := fmt.Sprintf("./%s/%s", r.baseDir, fileInfo.Name())
			file, err := os.Open(fileName)
			if err != nil {
				log.Println(err)
				continue
			}

			fileContent, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println(err)
				continue
			}

			var document model.Document
			if err := json.Unmarshal(fileContent, &document); err != nil {
				log.Println(err)
				continue
			}

			returnValue = append(returnValue, document)
		}

		r.refreshCache = false
	}

	return returnValue, nil
}

func (r *FileRepository) Save(document model.Document) error {
	if _, err := os.Stat(fmt.Sprintf("./%s", r.baseDir)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("./%s", r.baseDir), os.ModePerm)
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	documentBytes, err := json.Marshal(document)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fmt.Sprintf("./%s/%s.json", r.baseDir, document.Id), documentBytes, os.ModePerm); err != nil {
		return err
	}

	r.documents[document.Id] = document

	return nil
}

func (r *FileRepository) ById(id string) (model.Document, error) {
	document, exists := r.documents[id]

	if !exists {
		return document, ErrDocumentNotFound
	}

	return document, nil
}
