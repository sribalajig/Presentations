package mongo

import (
	"infra-balaji-rao/prezi.api.contracts/request"
	"log"
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
	log.Println("Reached the mongo persistence layer..")

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	results := reflect.New(typ).Interface()

	session := mongo.sessionFactory.Get()

	session.DB("prezi").C(mongo.collectionResolver.Resolve(typ)).Find(nil).All(results)

	return results
}
