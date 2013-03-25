package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/goredis"
)

type RedirectController struct {
	beego.Controller
}

func (this *RedirectController) Get() {
	this.Ctx.Request.ParseForm()
	short_url := this.Ctx.Request.Form.Get(":shorturl")
	var client goredis.Client
    val, _ := client.Get(short_url)
    if string(val)!="" {
        this.Ctx.Redirect(302, string(val))
    }else{
        this.Ctx.Redirect(302, "/static/html/404.html")
    }
}
