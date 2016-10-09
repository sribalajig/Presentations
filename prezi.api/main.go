package main

import (
	"github.com/astaxie/beego"
	_ "infra-balaji-rao/prezi.api/docs"
	_ "infra-balaji-rao/prezi.api/routers"
	"infra-balaji-rao/prezi.core/persistence/mongo"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	mongo.SessFactory = mongo.NewSessionFactory("localhost")

	beego.Run()
}
