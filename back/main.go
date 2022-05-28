package main

import (
	"beegoBackstage/controllers/ErrorControllers"
	"beegoBackstage/models"
	_ "beegoBackstage/routers"
	_ "beegoBackstage/utils"
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

func main() {
	logs.Debug(beego.BConfig.RunMode, "runMode")
	// 注册自定义错误
	beego.ErrorController(&ErrorControllers.ErrorController{})
	// 初始化数据库
	mysqlInit()

	beego.Run()
}
