package controllers

import (
	"github.com/astaxie/beego"
)

//MainController default page controller
type MainController struct {
	beego.Controller
}

//GetIndex get request.
func (c *MainController) GetIndex() {
	c.TplName = "index.html"
}

//GetLogin get request.
func (c *MainController) GetLogin() {
	c.TplName = "login.html"
}
