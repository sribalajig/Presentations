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

	presentations, _ := presentationService.Get(request.PresentationRequest{})

	presentationController.Data["json"] = presentations

	presentationController.ServeJSON()
}
