package main

import (
	_ "beegoBackstage/routers"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	val, _ := config.String("appname")
	logs.Info("auto load config name is", val)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
