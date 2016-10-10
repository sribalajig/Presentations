package routers

import (
	"infra-balaji-rao/prezi.web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
