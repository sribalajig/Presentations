package repository

import (
	"infra-balaji-rao/prezi.core/model"
)

type PresentationRepository struct {
}

func (presentationRepository PresentationRepository) Get() ([]model.Presentation, error) {
	return []model.Presentation{}, nil
}
