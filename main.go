package main

import (
	"issues/models"
	_ "issues/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetViewsPath("static")
	beego.InsertFilter("/*", beego.BeforeRouter, models.FilterLogin)
	beego.Run()
}
