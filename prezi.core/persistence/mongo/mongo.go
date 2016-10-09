package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"infra-balaji-rao/prezi.api.contracts/request"
	"reflect"
)

type Mongo struct {
	sessionFactory     SessionFactory
	collectionResolver CollectionResolver
}

var SessFactory SessionFactory

func NewMongo() Mongo {
	return Mongo{
		sessionFactory:     SessFactory,
		collectionResolver: CollectionResolver{},
	}
}

func (mongo Mongo) Get(typ reflect.Type, request request.Request) interface{} {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	results := reflect.New(typ).Interface()

	session := mongo.sessionFactory.Get()

	filteredRecords := session.DB("prezi").C(mongo.collectionResolver.Resolve(typ)).Find(nil)

	var sortedRecords, paginatedRecords *mgo.Query

	if request.SortingOption != nil {
		sortedRecords = filteredRecords.Sort(request.SortingOption.Field)
		paginatedRecords = sortedRecords.Skip(request.PaginationOption.Index).Limit(request.PaginationOption.NumberOfItems)
	} else if request.PaginationOption != nil {
		paginatedRecords = filteredRecords.Skip(request.PaginationOption.Index).Limit(request.PaginationOption.NumberOfItems)
	} else {
		paginatedRecords = filteredRecords
	}

	paginatedRecords.All(results)

	session.Close()

	return results
}
