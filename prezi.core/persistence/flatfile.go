package persistence

import (
	"infra-balaji-rao/prezi.core/model"
	"reflect"
	"time"
)

type FlatFile struct {
	path string
}

func NewFlatFile(path string) FlatFile {
	return FlatFile{
		path: path,
	}
}

func (flatFile FlatFile) Get(t reflect.Type) interface{} {
	return []model.Presentation{
		model.Presentation{
			Id:        "56f137f432fbb67217182785",
			Title:     "incididunt amet ad nostrud",
			ThumbNail: "https://placeimg.com/400/400/any",
			CreatedAt: time.Now(),
		},
	}
}
