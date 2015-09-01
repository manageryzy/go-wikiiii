package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/manageryzy/go-wikiiii/models"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const PERMISSION_EDIT = 1
const PERMISSION_EDIT_SCRIPT = 2

type EditController struct {
	beego.Controller
	uid        int
	permission int
}

func (this *EditController) Prepare() {
	uid := this.GetSession("uid")
	perm := this.GetSession("permission")
	if uid == nil || perm == nil {
		this.Abort("403")
	} else {
		this.uid = uid.(int)
		this.permission = perm.(int)

		if this.permission&PERMISSION_EDIT == 0 {
			this.Abort("403")
		}

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
		return
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 2 {
		content := this.GetString("content")

		script := false
		enableScript := this.GetString("EnableScript")
		if enableScript == "on" {
			script = true
		}

		if this.permission&PERMISSION_EDIT_SCRIPT == 0 {
			this.Abort("403")
			return
		}

		h := md5.New()
		h.Write([]byte(urls[1]))
		Md5 := hex.EncodeToString(h.Sum(nil))

		dir := "./data/" + Md5 + "/"
		filePath := dir + string(strconv.FormatInt(time.Now().Unix(), 10)) + ".md"
		os.MkdirAll(dir, 0777)

		if models.PageEdit(urls[1], content, this.uid, !script, filePath) {
			this.Ctx.Redirect(302, "/page/"+urls[1])
			return
		} else {
			this.TplNames = "editFail.tpl"
			return
		}

	} else {
		this.Abort("403")
		return
	}
}
