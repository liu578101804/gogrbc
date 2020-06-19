package main

import (
	_ "github.com/liu578101804/gorbac/routers"
	_ "github.com/liu578101804/gorbac/sysinit"
	_ "github.com/liu578101804/gorbac/filter"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
