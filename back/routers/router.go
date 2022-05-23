// Package routers @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beegoBackstage/controllers/BaseControllers"
	"beegoBackstage/middleware"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSInclude(&BaseControllers.TestController{}),
		beego.NSInclude(&BaseControllers.UserController{}),
		beego.NSInclude(&BaseControllers.RoleController{}),

		// 白名单路由
		beego.NSNamespace("/login", beego.NSInclude(&BaseControllers.LoginController{})),
	)

	whiteList := [...]string{"login"}
	// 登录校验
	filterUrl := "("
	for _, v := range whiteList {
		filterUrl += v + "|"
	}
	filterUrl += ")"
	logs.Info(filterUrl, "验证白名单")
	beego.InsertFilter("/api/login/"+filterUrl, beego.BeforeRouter, middleware.FilterUser)
	beego.AddNamespace(ns)
}
