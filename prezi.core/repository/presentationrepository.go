package repository

import (
	"infra-balaji-rao/prezi.core/model"
	"infra-balaji-rao/prezi.core/persistence"
	"reflect"
)

type PresentationRepository struct {
	dataAccessLayer persistence.DataAccessLayer
}

func NewPresentationRepository() PresentationRepository {
	return PresentationRepository{
		dataAccessLayer: persistence.NewFlatFile(
			"/Users/balaji/Documents/Go/src/infra-balaji-rao/prezi.core/persistence/presentations.json"),
	}
}

func (presentationRepository PresentationRepository) Get() ([]model.Presentation, error) {
	presentations := presentationRepository.dataAccessLayer.Get(
		reflect.TypeOf(model.Presentation{})).([]model.Presentation)

	return presentations, nil
}
