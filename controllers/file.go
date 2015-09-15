package controllers

import (
	"github.com/astaxie/beego"
	"github.com/manageryzy/go-wikiiii/models"
	"net/url"
	"strings"
)

type FileController struct {
	beego.Controller
	uid        int
	permission int
}

func (this *FileController) Prepare() {
	uid := this.GetSession("uid")
	perm := this.GetSession("permission")
	if uid == nil || perm == nil {
		//		this.Abort("403")
	} else {
		this.uid = uid.(int)
		this.permission = perm.(int)
	}
}

func (this *FileController) Get() {
	url, err := url.QueryUnescape(this.Ctx.Input.Request.URL.String())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) < 2 {
		this.Abort("403")
	}

	switch urls[1] {
	case "get":
		if len(urls) != 3 {
			this.Abort("403")
		}

		file := models.File{FileName: urls[2]}
		err := models.O.Read(&file)
		if err != nil {
			this.Abort("404")
		}

		if file.Cdn == 1 {
			this.Ctx.Redirect(302, file.Url)
			return
		} else {
			this.Ctx.Redirect(302, file.Path[1:])
		}

	case "history":
	case "delete":
	default:
		this.Abort("403")
	}
}
