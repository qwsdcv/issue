package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"issues/db"
	"log"
)

//IssueController handle /issue request
type IssueController struct {
	beego.Controller
}

//Get handle get request
func (c *IssueController) Get() {

}

//Post handle post request
func (c *IssueController) Post() {
	var bean db.Article
	err := json.Unmarshal(c.Ctx.Input.RequestBody, bean)
	if err != nil {
		c.Abort("500")
	}
	log.Printf("%s", c.Ctx.Input.URL())
}
