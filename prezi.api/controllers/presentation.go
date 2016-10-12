package controllers

import (
	"github.com/astaxie/beego"
	"infra-balaji-rao/prezi.api.contracts/request"
	"infra-balaji-rao/prezi.api.contracts/response"
	serviceLayer "infra-balaji-rao/prezi.core/service"
	"strings"
)

type PresentationController struct {
	beego.Controller
}

/*Get returns presentations based on the filters & options*/
func (presentationController *PresentationController) Get() {
	request := presentationController.generateRequest()
	service := serviceLayer.NewPresentationService()

	count := service.Count(request)
	numberOfItems := count

	if request.PaginationOption != nil {
		numberOfItems = request.PaginationOption.NumberOfItems
	}

	presentationController.Data["json"] = response.PaginatedResponse{
		Results:      service.Get(request),
		TotalRecords: count,
		TotalPages:   count / numberOfItems,
	}

	presentationController.ServeJSON()
}

/*generateRequest() parses the HTTP request for the relevant params*/
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
		direction := request.Asc

		if presentationController.GetString("direction") == "desc" {
			direction = request.Desc
		}

		presentationRequest.SortingOption = &request.SortingOption{
			Direction: direction,
			Field:     presentationController.GetString("sortby"),
		}
	}

	/*filter="title:*text*"*/
	if presentationController.GetString("filter") != "" {
		filterValue := presentationController.GetString("filter")

		presentationRequest.Filters = &[]request.Filter{
			request.Filter{
				Name:  strings.Split(filterValue, "_")[0],
				Value: strings.Split(filterValue, "_")[1],
			},
		}
	}

	return presentationRequest
}
