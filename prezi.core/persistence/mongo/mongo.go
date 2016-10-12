package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	contracts "infra-balaji-rao/prezi.api.contracts/request"
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

func (mongo Mongo) Get(typ reflect.Type, request contracts.Request) interface{} {
	session := mongo.sessionFactory.Get()

	collection := session.DB("prezi").C(mongo.collectionResolver.Resolve(typ))

	var filteredRecords *mgo.Query

	if request.Filters != nil {
		filterRegex := (*request.Filters)[0].Value.(string)

		filteredRecords = collection.Find(
			bson.M{(*request.Filters)[0].Name: bson.RegEx{".*" + filterRegex + ".*", "i"}})
	} else {
		filteredRecords = collection.Find(nil)
	}

	if request.SortingOption != nil {
		if request.SortingOption.Direction == contracts.Asc {
			filteredRecords = filteredRecords.Sort(request.SortingOption.Field)
		} else if request.SortingOption.Direction == contracts.Desc {
			filteredRecords = filteredRecords.Sort("-" + request.SortingOption.Field)
		}
	}

	if request.PaginationOption != nil {
		filteredRecords = filteredRecords.Skip(
			(request.PaginationOption.Index - 1) * request.PaginationOption.NumberOfItems).Limit(
			request.PaginationOption.NumberOfItems)
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	results := reflect.New(typ).Interface()

	filteredRecords.All(results)

	defer session.Close()

	return results
}

func (mongo Mongo) Count(typ reflect.Type, request contracts.Request) int {
	session := mongo.sessionFactory.Get()

	var filteredRecords *mgo.Query
	collection := session.DB("prezi").C(mongo.collectionResolver.Resolve(typ))

	if request.Filters != nil {
		filterRegex := (*request.Filters)[0].Value.(string)

		filteredRecords = collection.Find(
			bson.M{(*request.Filters)[0].Name: bson.RegEx{".*" + filterRegex + ".*", "i"}})
	} else {
		filteredRecords = collection.Find(nil)
	}

	count, _ := filteredRecords.Count()

	return count
}
