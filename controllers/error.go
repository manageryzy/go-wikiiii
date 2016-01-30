package controllers
import (
	"github.com/astaxie/beego"
)

type Error struct {
	beego.Controller
}

func (this *Error) Error403()  {
	this.TplName = "403.tpl"
}

func (this *Error) Error404()  {
	this.TplName = "404.tpl"
}

func (this *Error) Error500()  {
	this.TplName = "500.tpl"
}
