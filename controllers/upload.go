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
)

var FileUploadAllowType = []string{
	// audio
	"au", "snd", "mid", "rmi", "mp3", "aif", "aifc", "aiff", "m3u", "ra", "ram", "wav",
	// image
	"bmp", "cod", "gif", "ief", "jpe", "jpeg", "jpg", "jfif", "svg", "tif", "tiff", "ras", "cmx", "ico", "png", "pbm", "pgm", "ppm", "rgb", "xbm", "xpm", "xwd",
	// video
	"mp2", "mpa", "mpe", "mpeg", "mpg", "mpv2", "mov", "qt", "lsf", "lsx", "asf", "asr", "asx", "avi", "movie",
}

type UploadController struct {
	beego.Controller
	uid        int
	permission int
}

func (this *UploadController) Prepare() {
	uid := this.GetSession("uid")
	perm := this.GetSession("permission")
	if uid == nil || perm == nil {
		this.Abort("403")
	} else {
		this.uid = uid.(int)
		this.permission = perm.(int)

		if this.permission&PERMISSION_UPLOAD == 0 {
			this.Abort("403")
		}

		if this.uid == 0 {
			this.Abort("403")
		}
	}
}

func (this *UploadController) Get() {
	url, err := url.QueryUnescape(this.Ctx.Input.Request.URL.String())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 1 {

		this.Data["Title"] = "file upload"

		this.TplNames = "upload.tpl"
	} else {
		this.Abort("403")
	}
}

func (this *UploadController) Post() {
	url, err := url.QueryUnescape(this.Ctx.Input.Request.URL.String())

	if err != nil {
		this.Abort("500")
		return
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 1 {
		//检查基本上传权限
		if this.permission&PERMISSION_UPLOAD == 0 {
			this.Abort("403")
			return
		}

		//检查是否已经存在
		file := models.File{FileName: this.GetString("name")}
		err := models.O.Read(&file)
		if err == nil {
			this.Data["ERROR"] = "文件重复上传，请先删除"
			this.TplNames = "err.tpl"
			return
		}

		_, header, err := this.GetFile("file")

		fileName := header.Filename

		//获得后缀名
		t := strings.Split(fileName, ".")
		if len(t) < 2 {
			this.Abort("403")
		}
		ext := t[len(t)-1]

		//检查后缀名
		allow := false
		for _, e := range FileUploadAllowType {
			if strings.EqualFold(ext, e) {
				allow = true
			}
		}
		if !allow {
			this.Abort("403")
		}

		h := md5.New()
		h.Write([]byte(fileName))
		Md5 := hex.EncodeToString(h.Sum(nil))

		dir := "./uploads/" + strconv.Itoa(this.uid) + "/"
		filePath := dir + Md5 + "." + ext
		os.MkdirAll(dir, 0777)

		err = this.SaveToFile("file", filePath)
		if err != nil {
			this.Abort("500")
		}

		url := models.CDNGetURL("/" + strconv.Itoa(this.uid) + "/" + Md5 + "." + ext)

		file.Path = filePath
		file.Url = url
		file.Uid = this.uid
		file.Cdn = 0
		models.O.Insert(&file)

		if models.CDNUploadFile(filePath, "/"+strconv.Itoa(this.uid)+"/"+Md5+"."+ext) {
			file.Cdn = 1
			models.O.Update(&file)
			this.Ctx.Redirect(302, "/upload")
			return
		} else {
			this.TplNames = "err.tpl"
			return
		}

	} else {
		this.Abort("403")
		return
	}
}
