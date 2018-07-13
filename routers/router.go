package routers

import (
	"issues/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/issues/menu", &controllers.IssueController{}, "post:AddMenu")
}
