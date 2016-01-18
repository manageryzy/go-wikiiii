package controllers
import (
	"github.com/astaxie/beego"
)

type Error struct {
	beego.Controller
}

func (this *Error) Error403()  {
	this.TplNames = "403.tpl"
}

func (this *Error) Error404()  {
	this.TplNames = "404.tpl"
}

func (this *Error) Error500()  {
	this.TplNames = "500.tpl"
}
