package main

import (
	"beegoBackstage/models"
	_ "beegoBackstage/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接
func mysqlInit() {
	env, _ := beego.AppConfig.String("runmode")
	user, _ := beego.AppConfig.String("msql_username")
	passwd, _ := beego.AppConfig.String("msql_password")
	host, _ := beego.AppConfig.String("mysql_host")
	port, _ := beego.AppConfig.String("mysql_port")
	dbname, _ := beego.AppConfig.String("mysql_dbname")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", user, passwd, host, port, dbname)
	logs.Debug(dataSource)
	orm.RegisterDataBase("default", "mysql", dataSource)
	// 注册model
	models.RegisterModels()

	if env == "dev" {
		// 开启查询调试模式
		orm.Debug = true
	}
}

func init() {
	mysqlInit()
}

func main() {
	logs.Debug(beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
