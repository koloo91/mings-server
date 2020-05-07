package repository

import "github.com/koloo91/mings-server/model"

type Repository interface {
	Get() ([]model.Document, error)
	Save(document model.Document) error
	ById(id string) (model.Document, error)
}
