package controllers

import (
	"encoding/json"
	"issues/models"

	"github.com/astaxie/beego"

	"log"
)

//IssueController handle /issue request
type IssueController struct {
	beego.Controller
}

//Get handle get request
func (c *IssueController) Get() {

}

//LoadMenu handle get request and return Menu info as JSON
func (c *IssueController) LoadMenu() {
	obj, err := models.GetMenu()
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = obj
	c.ServeJSON()
}

//AddMenu handle post request and return Menu info as JSON
func (c *IssueController) AddMenu() {
	log.Println(c.Ctx.Input.RequestBody)

	var bean models.Article
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &bean)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	log.Printf("%s", c.Ctx.Input.URL())

	bean.AddMenu()
	obj, err := models.GetMenu()
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = obj

	c.ServeJSON()
}
