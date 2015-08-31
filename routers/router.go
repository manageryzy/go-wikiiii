package routers

import (
	"github.com/manageryzy/go-wikiiii/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	
	beego.Router("/page/*", &controllers.PageController{})
	beego.Router("/category/*", &controllers.CategoryController{})
	
	beego.Router("/edit/*", &controllers.EditController{})
	
	beego.Router("/login", &controllers.LoginController{})
}
