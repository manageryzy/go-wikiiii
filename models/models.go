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

type Category struct {
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

var O orm.Ormer

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Page))
}
