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
	Title    string `orm:"pk;index"`
	Category string `orm:"index"`
}

type Page struct {
	Title    string `orm:"pk;index"`
	Page     string `orm:"type(text)"`
	Uid      int
	Lastedit time.Time `orm:"auto_now;index;type(datetime)"`
}

type File struct {
	FileName string `orm:"pk;index"`
	Path     string `orm:"type(text)"`
	Url      string `orm:"type(text)"`
}

type History struct {
	Title  string    `orm:"pk;index"`
	Path   string    `orm:"type(text)"`
	Uid    int       `orm:"index"`
	Update time.Time `orm:"auto_now;index;type(datetime)"`
}

var O orm.Ormer

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Page), new(Categories), new(File), new(History))
}
