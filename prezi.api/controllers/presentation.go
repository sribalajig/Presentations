package controllers

import (
	"github.com/astaxie/beego"
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.api.contracts/response"
	"infra-balaji-rao/prezi.core/service"
)

type PresentationController struct {
	beego.Controller
}

func (presentationController *PresentationController) Get() {
	presentations, _ := service.NewPresentationService().Get(
		presentationController.generateRequest())

	presentationController.Data["json"] = response.PaginatedResponse{
		Results:      presentations,
		TotalRecords: 100,
		TotalPages:   10,
	}

	presentationController.ServeJSON()
}

func (presentationController *PresentationController) generateRequest() request.Request {
	presentationRequest := request.Request{}

	paginate, _ := presentationController.GetBool("paginate")

	if paginate {
		index, _ := presentationController.GetInt("index")
		numberOfItems, _ := presentationController.GetInt("numitems")

		presentationRequest.PaginationOption = &request.PaginationOption{
			Index:         index,
			NumberOfItems: numberOfItems,
		}
	}

	sort, _ := presentationController.GetBool("sort")

	if sort {
		direction, _ := presentationController.GetInt("direction")

		presentationRequest.SortingOption = &request.SortingOption{
			Direction: direction,
			Field:     presentationController.GetString("sortby"),
		}
	}

	return presentationRequest
}
