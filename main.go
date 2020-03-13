package main

import (
	_ "ai-config-project/routers"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/ai_config?charset=utf8")

}
func main() {
	o := orm.NewOrm()
	o.Using("default")
	beego.Run()
}
