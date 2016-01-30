package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/manageryzy/go-wikiiii/models"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	url, err := url.QueryUnescape(this.Ctx.Input.URL())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 2 {
		var maps []orm.Params
		num, err := models.O.QueryTable("history").Filter("title", urls[1]).OrderBy("-update").Limit(1).Values(&maps)
		if err != nil {
			this.Abort("500")
		}

		if num != 0 {
			b, e := ioutil.ReadFile(maps[0]["Path"].(string))

			if e != nil {
				this.Abort("500")
			}
			this.Data["Src"] = string(b)
		} else {
			this.Data["Src"] = ""
		}

		this.Data["Title"] = urls[1]

		this.TplName = "edit.tpl"
	} else {
		this.Abort("403")
	}
}

func (this *EditController) Post() {
	url, err := url.QueryUnescape(this.Ctx.Input.URL())

	if err != nil {
		this.Abort("500")
		return
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 2 {
		content := this.GetString("content")
		reason := this.GetString("reason")

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

		if models.PageEdit(urls[1], content, this.uid, !script, filePath, reason) {
			this.Ctx.Redirect(302, "/page/"+urls[1])
			return
		} else {
			this.TplName = "err.tpl"
			return
		}

	} else {
		this.Abort("403")
		return
	}
}
