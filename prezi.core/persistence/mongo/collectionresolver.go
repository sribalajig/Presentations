package mongo

import (
	"infra-balaji-rao/prezi.core/model"
	"log"
	"reflect"
)

type CollectionResolver struct {
}

func (collectionResolver CollectionResolver) Resolve(typ reflect.Type) string {
	if typ == reflect.TypeOf([]model.Presentation{}) {
		log.Println("Need to use the presentation collection.")

		return "presentation"
	}

	log.Println(typ)
	log.Println(model.Presentation{})

	log.Println("Need to use the default collection.")

	return ""
}
