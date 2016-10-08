package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:PresentationController"] = append(beego.GlobalControllerRouter["infra-balaji-rao/prezi.api/controllers:PresentationController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

}
