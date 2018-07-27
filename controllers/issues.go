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

//LoadContent handle get request and return content info as JSON
func (c *IssueController) LoadContent() {
	id := c.Ctx.Input.Param(":id")
	obj, err := models.GetContent(id)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = obj

	c.ServeJSON()
}

//SetContent set the content of the article ${id}.
func (c *IssueController) SetContent() {
	id := c.Ctx.Input.Param(":id")
	var bean models.Article
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &bean)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	err = models.SetContent(id, bean.Content)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body([]byte("{}"))
}

//Login enter root mode, which means you can modify the contents.
func (c *IssueController) Login() {
	var secret login
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &secret)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	if secret.Secret == beego.AppConfig.String("secret") {

		token, err := models.CreateToken()
		if err != nil {
			c.CustomAbort(500, err.Error())
		}
		c.Ctx.Output.Header("Authorization", token)
		c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		c.Ctx.Output.Body([]byte("{}"))
	} else {
		c.CustomAbort(401, "You cannot pass.")
	}
}

type login struct {
	Secret string `json:"secret"`
}
