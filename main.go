package main

import (
	_ "issues/models"
	_ "issues/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
