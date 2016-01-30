package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/manageryzy/go-wikiiii/models"
	"net/url"
	"strconv"
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
		this.uid = 0
		this.permission = 0
	} else {
		this.uid = uid.(int)
		this.permission = perm.(int)
	}
}

func (this *FileController) Get() {
	url, err := url.QueryUnescape(this.Ctx.Input.URL())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) < 2 {
		this.Abort("404")
	}

	switch urls[1] {
	case "get":
		if len(urls) != 3 {
			this.Abort("404")
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
			return
		}

	case "history":
		if len(urls) < 3 {
			this.Data["ERROR"] = "参数错误"
			this.TplName = "err.tpl"
			return
		}

		switch urls[2] {
		case "list": //列出历史记录
			var maps []orm.Params

			name := this.GetString("name", "")
			if name == "" {
				if this.uid == 0 {
					this.Abort("403")
				}

				if this.permission&PERMISSION_VIEW_UPLOAD != 0 {
					models.O.QueryTable("history_file").OrderBy("update").Limit(50).Values(&maps)
				} else {
					models.O.QueryTable("history_file").Filter("uid", this.uid).OrderBy("update").Limit(50).Values(&maps)
				}
			} else {
				if this.permission&PERMISSION_VIEW_UPLOAD != 0 {
					models.O.QueryTable("history_file").Filter("file_name", name).OrderBy("update").Values(&maps)
				} else {
					models.O.QueryTable("history_file").Filter("file_name", name).Filter("uid", this.uid).OrderBy("update").Values(&maps)
				}
			}

			this.Data["History"] = maps
			this.TplName = "history_file.tpl"

		case "get": //获得旧版文件
			if len(urls) < 4 {
				this.Data["ERROR"] = "参数错误"
				this.TplName = "err.tpl"
				return
			}

			fhid, _ := strconv.Atoi(urls[3])
			file := models.HistoryFile{Fhid: fhid}
			err := models.O.Read(&file)
			if err != nil {
				this.Abort("404")
			}

			if file.Cdn == 1 {
				this.Ctx.Redirect(302, file.Url)
				return
			} else {
				this.Ctx.Redirect(302, file.Path[1:])
				return
			}

		case "view": //旧版文件信息
			if len(urls) < 4 {
				this.Data["ERROR"] = "参数错误"
				this.TplName = "err.tpl"
				return
			}

			fhid, _ := strconv.Atoi(urls[3])
			file := models.HistoryFile{Fhid: fhid}
			err := models.O.Read(&file)
			if err != nil {
				this.Abort("404")
			}
		case "use": //使用旧版文件代替新版文件
			if len(urls) < 4 {
				this.Data["ERROR"] = "参数错误"
				this.TplName = "err.tpl"
				return
			}
			fhid, _ := strconv.Atoi(urls[3])
			file := models.HistoryFile{Fhid: fhid}
			err := models.O.Read(&file)
			if err != nil {
				this.Abort("404")
			}

			if this.permission&PERMISSION_EDIT_UPLOAD == 0 {
				this.Abort("403")
			}

			f := models.File{Uid: file.Uid, FileName: file.FileName, Path: file.Path, Url: file.Url, Cdn: file.Cdn}
			models.O.Update(&f)
		default:
			this.Abort("404")
		}
	case "delete":
	default:
		this.CustomAbort(404,"404")
	}
}
