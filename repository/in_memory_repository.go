package repository

import (
	"github.com/koloo91/mings-server/model"
	"sync"
)

type InMemoryRepository struct {
	documents map[string]model.Document
	mutex     sync.Mutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		documents: make(map[string]model.Document),
	}
}

func (r *InMemoryRepository) Get() ([]model.Document, error) {
	returnValue := make([]model.Document, 0, len(r.documents))

	for _, document := range r.documents {
		returnValue = append(returnValue, document)
	}

	return returnValue, nil
}

func (r *InMemoryRepository) Save(document model.Document) error {
	r.mutex.Lock()
	r.documents[document.Id] = document
	r.mutex.Unlock()
	return nil
}

func (r *InMemoryRepository) ById(id string) (model.Document, error) {
	document, exists := r.documents[id]

	if !exists {
		return document, ErrDocumentNotFound
	}

	return document, nil
}

func (r *InMemoryRepository) GetByDependsOn(dependsOn string) ([]model.Document, error) {
	panic("implement me")
}
