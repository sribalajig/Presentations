package repository

import (
	"infra-balaji-rao/prezi.core/persistence"
	"reflect"
)

type Repository struct {
	typ             reflect.Type
	dataAccessLayer persistence.DataAccessLayer
}

func NewRepository(typ reflect.Type) Repository {
	return Repository{
		typ: typ,
		dataAccessLayer: persistence.NewFlatFile(
			"/Users/balaji/Documents/Go/src/infra-balaji-rao/prezi.core/persistence/presentations.json"),
	}
}

func (repository Repository) Get() (interface{}, error) {
	return repository.dataAccessLayer.Get(repository.typ), nil
}
