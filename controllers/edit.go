package controllers

import (
	"github.com/astaxie/beego"
	"github.com/manageryzy/go-wikiiii/models"
	"net/url"
	"strings"
)

type EditController struct {
	beego.Controller
	uid int
}

func (this *EditController) Prepare() {
	uid := this.GetSession("uid")
	if uid == nil {
		this.Abort("403")
	} else {
		this.uid = uid.(int)
		if this.uid == 0 {
			this.Abort("403")
		}
	}
}

func (this *EditController) Get() {
	url, err := url.QueryUnescape(this.Ctx.Input.Request.URL.String())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 2 {
		page, exist := models.PageGetSQL(urls[1])
		if exist {
			this.Data["Src"] = page
		} else {
			this.Data["Src"] = ""
		}

		this.Data["Title"] = urls[1]

		this.TplNames = "edit.tpl"
	} else {
		this.Abort("403")
	}
}

func (this *EditController) Post() {
	url, err := url.QueryUnescape(this.Ctx.Input.Request.URL.String())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 2 {
		content := this.GetString("content")

		if models.PageEdit(urls[1], content, this.uid) {
			this.Ctx.Redirect(302, "/page/"+urls[1])
		} else {
			this.TplNames = "editFail.tpl"
		}

	} else {
		this.Abort("403")
	}
}
