package main

import (
	_ "nepliteApi/routers"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego"
	_ "nepliteApi/models"
	"github.com/astaxie/beego/orm"
)
func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "neplite.db")
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
