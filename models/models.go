package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Uid        int `orm:"pk;index"`
	Name       string
	Permission int
}

type Categories struct {
	Id       int    `orm:"pk;auto"`
	Title    string `orm:"index"`
	Category string `orm:"index"`
}

type Page struct {
	Title    string `orm:"pk;index"`
	Page     string `orm:"type(text)"`
	Uid      int
	Lastedit time.Time `orm:"auto_now;index;type(datetime)"`
	Safe     int
}

type File struct {
	FileName string `orm:"pk;index"`
	Path     string `orm:"type(text)"`
	Url      string `orm:"type(text)"`
	Uid      int
	Cdn      int
}

type History struct {
	Hid    int    `orm:"pk"`
	Title  string `orm:"index"`
	Path   string `orm:"type(text)"`
	Reason string `orm:"type(text)"`
	Uid    int    `orm:"index"`
	Name   string
	Update time.Time `orm:"auto_now;index;type(datetime)"`
}

type HistoryFile struct {
	Fhid     int    `orm:"pk"`
	FileName string `orm:"index"`
	Path     string `orm:"type(text)"`
	Url      string `orm:"type(text)"`
	Uid      int
	Name     string
	Update   time.Time `orm:"auto_now;index;type(datetime)"`
	Cdn      int
}

var O orm.Ormer

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Page), new(Categories), new(File), new(History), new(HistoryFile))
}
