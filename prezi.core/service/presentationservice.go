package service

import (
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.core/model"
)

type PresentationService struct {
}

func (presentationService PresentationService) Get(request request.PresentationRequest) ([]model.Presentation, error) {
	return []model.Presentation{}, nil
}
