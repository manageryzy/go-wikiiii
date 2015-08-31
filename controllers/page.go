package controllers

import (
	"github.com/astaxie/beego"
	"github.com/manageryzy/go-wikiiii/models"
	"strings"
)

type PageController struct {
	beego.Controller
}

func (this *PageController) Get() {
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"


	url := strings.Trim(this.Ctx.Input.Request.URL.String(),"/")
	urls := strings.Split(url,"/")

	if len(urls) == 1{
		this.Abort("404")
	}else if len(urls) == 2{
		this.Data["Title"] = urls[1]
		this.Data["Page"] = models.PageGet(urls[1])
		this.TplNames = "page.tpl"
	}else if len(urls) == 3 {
		if urls[2] == "edit"{
			this.TplNames = "edit.tpl"
		}else {
			this.Abort("403")
		}
	}else{
		this.Abort("403")
	}
}
