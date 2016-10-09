package service

import (
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.core/model"
	"infra-balaji-rao/prezi.core/repository"
	"reflect"
)

type PresentationService struct {
	presentationRepository repository.Repository
}

func NewPresentationService() PresentationService {
	return PresentationService{
		presentationRepository: repository.NewRepository(reflect.TypeOf(&[]model.Presentation{})),
	}
}

func (presentationService PresentationService) Get(request request.Request) (*[]model.Presentation, error) {
	presentations, nil := presentationService.presentationRepository.Get(request)

	return presentations.(*[]model.Presentation), nil
}
