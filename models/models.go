package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Uid         int 		`orm:"pk;index"`
	Name        string
	Permission	int
}

type Page struct  {
	Title 		string		`orm:"pk;index"`
	Page		string		`orm:"type(text)"`
	Lastedit	time.Time 	`orm:"auto_now;index;type(datetime)"`
}

var O orm.Ormer

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User),new(Page))
}
