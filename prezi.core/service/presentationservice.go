package service

import (
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.core/model"
	"time"
)

type PresentationService struct {
}

func (presentationService PresentationService) Get(request request.PresentationRequest) ([]model.Presentation, error) {
	return []model.Presentation{
		model.Presentation{
			Id:        "56f137f432fbb67217182785",
			Title:     "incididunt amet ad nostrud",
			ThumbNail: "https://placeimg.com/400/400/any",
			CreatedAt: time.Now(),
		},
	}, nil
}
