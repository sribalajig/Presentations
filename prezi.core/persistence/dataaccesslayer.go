package persistence

import (
	"infra-balaji-rao/prezi.api.contracts/request"
	"reflect"
)

type DataAccessLayer interface {
	Get(t reflect.Type, request request.Request) interface{}
}
