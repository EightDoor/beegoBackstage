// Package routers @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	DemoCrud2 "beegoBackstage/controllers/DemoCrud"
	"beegoBackstage/controllers/JournalControllers"
	"beegoBackstage/controllers/SysControllers"
	"beegoBackstage/controllers/UploadControllers"
	"beegoBackstage/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			// sys
			beego.NSInclude(&SysControllers.TestController{}),
			beego.NSInclude(&SysControllers.UserController{}),
			beego.NSInclude(&SysControllers.RoleController{}),
			beego.NSInclude(&SysControllers.MenuController{}),
			beego.NSInclude(&SysControllers.DeptController{}),
			beego.NSInclude(&SysControllers.DictController{}),
			beego.NSInclude(&SysControllers.DictItemController{}),
			beego.NSInclude(&SysControllers.SysUserRoleController{}),
			beego.NSInclude(&SysControllers.SysRoleMenuController{}),

			// user log
			beego.NSInclude(&JournalControllers.LoginLogController{}),
			beego.NSInclude(&JournalControllers.LogRequestController{}),

			// upload
			beego.NSInclude(&UploadControllers.UploadController{}),

			// other
			beego.NSInclude(&DemoCrud2.DemoCrudController{}),

			// 白名单路由
			beego.NSNamespace("/login", beego.NSInclude(&SysControllers.LoginController{})),
		),
	)
	// 登录校验
	beego.InsertFilter("/api/v1/*", beego.BeforeRouter, middleware.FilterUser)
	beego.AddNamespace(ns)
}
