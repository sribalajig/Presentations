package persistence

import (
	"encoding/json"
	"fmt"
	"infra-balaji-rao/prezi.api.contracts/request"
	"io/ioutil"
	"log"
	"reflect"
)

type FlatFile struct {
	path string
}

func NewFlatFile(path string) FlatFile {
	return FlatFile{
		path: path,
	}
}

func (flatFile FlatFile) Get(typ reflect.Type, request request.Request) interface{} {
	data, _ := ioutil.ReadFile(flatFile.path)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	payload := reflect.New(typ).Interface()

	if err := json.Unmarshal(data, payload); err != nil {
		log.Println("Error unmarshalling json : " + err.Error())

		return nil
	} else {
		log.Println(fmt.Sprintf("This is the unmarshalled data %q", payload))
	}

	return payload
}
