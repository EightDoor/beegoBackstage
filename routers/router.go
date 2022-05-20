// Package routers @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beegoBackstage/controllers"
	"beegoBackstage/middleware"
	_ "beegoBackstage/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/api"),
		beego.NSNamespace("/login", beego.NSInclude(&controllers.LoginController{})),
	)

	// 登录校验
	beego.InsertFilter("/v1/api/*", beego.BeforeRouter, middleware.FilterUser)
	beego.AddNamespace(ns)
}
