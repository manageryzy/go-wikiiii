package controllers

import (
	"github.com/astaxie/beego"
	"github.com/manageryzy/go-wikiiii/models"
	"net/url"
	"strings"
)

type PageController struct {
	beego.Controller
}

func (this *PageController) Get() {
	url, err := url.QueryUnescape(this.Ctx.Input.Request.URL.String())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 1 {
		this.Abort("404")
	} else if len(urls) == 2 {
		//普通页面
		this.Data["Title"] = urls[1]
		this.Data["Page"] = models.PageGet(urls[1])
		this.Data["Category"] = models.PageGetCategory(urls[1])
		this.TplNames = "page.tpl"
	} else if len(urls) == 3 {
		if urls[2] == "category" {
			//分类目录
			this.Data["Title"] = urls[1]
			this.Data["Page"] = models.PageGet(urls[1])
			this.Data["Category"] = models.PageGetCategory(urls[1])
			this.Data["Pages"] = models.CategoryGetPages(urls[1])
			this.TplNames = "category.tpl"
		} else if urls[2] == "history" {
			//历史列表
			this.TplNames = "history.tpl"
		} else {
			this.Abort("403")
		}
	} else if len(urls) == 4 {
		//历史页面
		this.Data["Title"] = urls[1] + " history "
		this.Data["Page"] = models.PageGet(urls[1])
		this.Data["Category"] = models.PageGetCategory(urls[1])
		this.TplNames = "page.tpl"
	} else {
		this.Abort("403")
	}
}
