package main

import (
	_ "nepliteApi/routers"
	//_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego"
	_ "nepliteApi/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/logs"
)

func init() {
	dbType := beego.AppConfig.String("db_type")
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDriver(dbType, orm.DRMySQL)
	//orm.RegisterDataBase("default", "sqlite3", "neplite.db")
	// dbCharset := beego.AppConfig.String(dbType + "db_charset")
	dataSource := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	logs.Info("%s==== %s", dbAlias, dataSource)
	orm.RegisterDataBase(dbAlias, dbType, dataSource)
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
