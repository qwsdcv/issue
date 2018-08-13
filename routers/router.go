package routers

import (
	"issues/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:GetIndex")
	beego.Router("/index.html", &controllers.MainController{}, "get:GetIndex")
	beego.Router("/issues/menu", &controllers.IssueController{}, "post:AddMenu")
	beego.Router("/issues/menu", &controllers.IssueController{}, "get:LoadMenu")
	beego.Router("/issues/content/:id", &controllers.IssueController{}, "get:LoadContent")
	beego.Router("/issues/content/:id", &controllers.IssueController{}, "post:SetContent")

	beego.Router("/issues/comment/:id", &controllers.IssueController{}, "get:GetComment")
	beego.Router("/issues/comment", &controllers.IssueController{}, "post:AddComment")

	beego.Router("/issues/secret", &controllers.IssueController{}, "post:Login")
}
