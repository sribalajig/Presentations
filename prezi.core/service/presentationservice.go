package service

import (
	"infra-balaji-rao/prezi/model"
)

type PresentationService struct {
}

func (presentationService PresentationService) Get() ([]model.Presentation, error) {
	return []model.Presentation{}, nil
}
