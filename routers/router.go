package routers

import (
	"github.com/astaxie/beego"
	"github.com/manageryzy/go-wikiiii/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/page/*", &controllers.PageController{})
	beego.Router("/category/*", &controllers.CategoryController{})

	beego.Router("/edit/*", &controllers.EditController{})
	beego.Router("/file/*", &controllers.FileController{})
	beego.Router("/upload", &controllers.UploadController{})

	beego.SetStaticPath("/uploads", "uploads")

	beego.Router("/login", &controllers.LoginController{})
}
