package service

import (
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.core/model"
	"infra-balaji-rao/prezi.core/repository"
)

type PresentationService struct {
	presentationRepository repository.PresentationRepository
}

func NewPresentationService() PresentationService {
	return PresentationService{
		presentationRepository: repository.NewPresentationRepository(),
	}
}

func (presentationService PresentationService) Get(request request.PresentationRequest) ([]model.Presentation, error) {
	presentations, _ := presentationService.presentationRepository.Get()

	return presentations, nil
}
