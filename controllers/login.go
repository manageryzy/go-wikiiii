package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//this.Data["Website"] = "beego.me"

	this.SetSession("uid", 1)
	this.SetSession("permission", PERMISSION_EDIT|PERMISSION_EDIT_SCRIPT|PERMISSION_UPLOAD|PERMISSION_VIEW_UPLOAD|PERMISSION_EDIT_UPLOAD)
	this.TplNames = "login.tpl"
}
