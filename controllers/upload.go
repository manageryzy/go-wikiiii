package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/manageryzy/go-wikiiii/models"
	"io/ioutil"
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
	url, err := url.QueryUnescape(this.Ctx.Input.URL())

	if err != nil {
		this.Abort("500")
	}

	url = strings.Trim(url, "/")
	urls := strings.Split(url, "/")

	if len(urls) == 1 {

		this.Data["Title"] = "file upload"

		this.TplName = "upload.tpl"
	} else {
		this.Abort("403")
	}
}

func (this *UploadController) Post() {
	url, err := url.QueryUnescape(this.Ctx.Input.URL())

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
		fileExist := false
		file := models.File{FileName: this.GetString("name")}
		err := models.O.Read(&file)
		if err == nil {
			fileExist = true
		}

		_, header, err := this.GetFile("file")
		if err != nil {
			this.Data["ERROR"] = err.Error()
			this.TplName = "err.tpl"
			return
		}

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

		//确定临时文件文件名
		h := md5.New()
		h.Write([]byte(fileName))
		Md5 := hex.EncodeToString(h.Sum(nil))

		//保存临时文件
		tmpFileName := "/tmp/" + Md5
		err = this.SaveToFile("file", tmpFileName)
		if err != nil {
			this.Abort("500")
		}

		//读取临时文件并且计算md5
		content, err := ioutil.ReadFile(tmpFileName)
		if err != nil {
			this.Data["error"] = "写入临时文件失败"
			this.TplName = "err.tpl"
		}
		h.Reset()
		h.Write(content)
		Md5 = hex.EncodeToString(h.Sum(nil))

		dir := "./uploads/" + strconv.Itoa(this.uid) + "/"
		filePath := dir + Md5 + "." + ext
		os.MkdirAll(dir, 0777)

		err = this.SaveToFile("file", filePath)
		if err != nil {
			this.Abort("500")
		}

		url := models.CDNGetURL("/" + strconv.Itoa(this.uid) + "/" + Md5 + "." + ext)

		user := models.User{Uid: this.uid}
		err = models.O.Read(&user)
		if err != nil {
			this.Abort("403")
		}

		file.Path = filePath
		file.Url = url
		file.Uid = this.uid
		file.Cdn = 0
		if !fileExist {
			models.O.Insert(&file)
		} else {
			models.O.Update(&file)
		}

		if models.CDNUploadFile(filePath, "/"+strconv.Itoa(this.uid)+"/"+Md5+"."+ext) {
			file.Cdn = 1
			history := models.HistoryFile{FileName: this.GetString("name"), Path: filePath, Url: url, Uid: this.uid, Cdn: 1, Name: user.Name}
			models.O.Insert(&history)
			models.O.Update(&file)
			this.Ctx.Redirect(302, "/upload")
			return
		} else {
			history := models.HistoryFile{FileName: this.GetString("name"), Path: filePath, Url: url, Uid: this.uid, Cdn: 0, Name: user.Name}
			models.O.Insert(&history)
			this.TplName = "err.tpl"
			return
		}

	} else {
		this.Abort("403")
		return
	}
}
