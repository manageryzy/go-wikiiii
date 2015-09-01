package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manageryzy/go-wikiiii/models"
	_ "github.com/manageryzy/go-wikiiii/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/"+beego.AppConfig.String("mysqldb")+"?charset=utf8")
	orm.Debug = true
}

func main() {
	models.O = orm.NewOrm()
	models.O.Using("default") // 默认使用 default，你可以指定为其他数据库

	var err error
	models.PageCache, err = cache.NewCache("memory", `{"interval":60}`)

	if err != nil {
		println(err.Error())
		return
	}

	beego.Run()
}
