package persistence

import (
	"reflect"
)

type DataAccessLayer interface {
	Get(t reflect.Type) interface{}
}
