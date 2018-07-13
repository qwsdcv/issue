package controllers

import (
	"github.com/astaxie/beego"
)

//MainController default page controller
type MainController struct {
	beego.Controller
}

//Get get request.
func (c *MainController) Get() {
	c.TplName = "index.html"
}
