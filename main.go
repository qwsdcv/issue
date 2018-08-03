package main

import (
	"issues/models"
	_ "issues/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/go", "views")
	beego.InsertFilter("/*", beego.BeforeRouter, models.FilterLogin)
	beego.Run()
}
