package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:LoginController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:LoginController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
