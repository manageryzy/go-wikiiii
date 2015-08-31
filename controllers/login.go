package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//this.Data["Website"] = "beego.me"
	
	this.SetSession("uid",1)
	this.TplNames = "login.tpl"
}
