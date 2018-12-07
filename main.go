package main

import (
	"issues/models"
	_ "issues/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetViewsPath("static")
	beego.SetStaticPath("/js", "static/js")
	beego.InsertFilter("/*", beego.BeforeRouter, models.FilterDDOS)
	beego.InsertFilter("/*", beego.BeforeRouter, models.FilterLogin)
	beego.Run()
}
