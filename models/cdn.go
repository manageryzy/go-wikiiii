package models

import (
	"code.google.com/p/ftp4go"
	"github.com/astaxie/beego"
)

// 访问的时候CDN的URL，不包含目录的前缀
var CDN_ADDR = ""

// CDN类型，暂时只支持FTP上传
var CDN_TYPE = ""

// 在CDN上面储存的相对路径部分
var CDN_DIR = ""

// FTP config
var CDN_FTP_IP = ""
var CDN_FTP_PORT = 21
var CDN_FTP_USER = ""
var CDN_FTP_PWD = ""

func CDNInit() {
	CDN_ADDR = beego.AppConfig.String("cdn_addr")
	CDN_TYPE = beego.AppConfig.String("cdn_type")
	CDN_DIR = beego.AppConfig.String("cdn_dir")

	CDN_FTP_IP = beego.AppConfig.String("cdn_ftp_ip")
	CDN_FTP_PORT, _ = beego.AppConfig.Int("cdn_ftp_port")
	CDN_FTP_USER = beego.AppConfig.String("cdn_ftp_usre")
	CDN_FTP_PWD = beego.AppConfig.String("cdn_ftp_pwd")

}

func CDNGetURL(path string) (url string) {
	url = CDN_ADDR + CDN_DIR + path
	return
}

func CDNUploadFile(local string, remote string) bool {
	ftpClient := ftp4go.NewFTP(1)

	//connect
	_, err := ftpClient.Connect(CDN_FTP_IP, CDN_FTP_PORT, "")
	if err != nil {
		return false
	}

	defer ftpClient.Quit()

	_, err = ftpClient.Login(CDN_FTP_USER, CDN_FTP_PWD, "")
	if err != nil {
		return false
	}

	err = ftpClient.UploadFile(CDN_DIR+remote, local, false, func(info *ftp4go.CallbackInfo) {
		println("callback")
	})

	if err != nil {
		println(err.Error())
		return false
	}

	return true
}

func CDNRemoveFile(remote string) bool {
	return true
}
