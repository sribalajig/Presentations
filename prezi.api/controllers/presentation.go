package controllers

import (
	"github.com/astaxie/beego"
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.core/service"
)

type PresentationController struct {
	beego.Controller
}

func (presentationController *PresentationController) Get() {
	presentationService := service.NewPresentationService()

	presentations, _ := presentationService.Get(request.Request{
		PaginationOption: &request.PaginationOption{
			Index:         25,
			NumberOfItems: 35,
		},
		SortingOption: &request.SortingOption{
			Field:     "title",
			Direction: request.Asc,
		},
	})

	presentationController.Data["json"] = presentations

	presentationController.ServeJSON()
}
