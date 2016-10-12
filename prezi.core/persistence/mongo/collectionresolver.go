package mongo

import (
	"infra-balaji-rao/prezi.core/model"
	"reflect"
)

type CollectionResolver struct {
}

func (collectionResolver CollectionResolver) Resolve(typ reflect.Type) string {
	if typ == reflect.TypeOf(&[]model.Presentation{}) {
		return "presentation"
	}

	return ""
}
