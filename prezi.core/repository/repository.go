package repository

import (
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.core/persistence"
	"infra-balaji-rao/prezi.core/persistence/mongo"
	"reflect"
)

type Repository struct {
	typ             reflect.Type
	dataAccessLayer persistence.DataAccessLayer
}

func NewRepository(typ reflect.Type) Repository {
	return Repository{
		typ:             typ,
		dataAccessLayer: mongo.NewMongo(),
	}
}

func (repository Repository) Get(request request.Request) (interface{}, error) {
	return repository.dataAccessLayer.Get(repository.typ, request), nil
}
