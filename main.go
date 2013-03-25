package main

import (
	"shorturl/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.PprofOn = true
	beego.RegisterController("/", &controllers.MainController{})
	beego.RegisterController("/:shorturl([0-9a-zA-Z]{6})", &controllers.RedirectController{})
	beego.Run()
}
