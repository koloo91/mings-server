package service

import (
	"github.com/koloo91/mings-server/model"
	"github.com/koloo91/mings-server/repository"
)

var Repository repository.Repository

func GetAll() ([]model.Document, error) {
	return Repository.Get()
}

func Create(document model.Document) error {
	return Repository.Save(document)
}

func GetById(id string) (model.Document, error) {
	return Repository.ById(id)
}
