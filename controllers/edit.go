package controllers

import (
	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Prepare() {
	uid := this.GetSession("uid")
	if uid == nil || uid == 0{
		this.Abort("403")
	}
}

func (this *EditController) Get() {
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "edit.tpl"
}

func (this *EditController) Post(){
	
}