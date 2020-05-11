package repository

import (
	"errors"
	"github.com/koloo91/mings-server/model"
)

var (
	ErrDocumentNotFound = errors.New("document not found")
)

type Repository interface {
	Get() ([]model.Document, error)
	Save(document model.Document) error
	ById(id string) (model.Document, error)
}
