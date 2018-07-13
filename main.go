package main

import (
	_ "issues/db"
	_ "issues/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
