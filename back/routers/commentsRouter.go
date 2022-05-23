package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:LoginController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:LoginController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/role`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/role`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/role`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/role/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:RoleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/role/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/test`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:TestController"],
        beego.ControllerComments{
            Method: "GetPaging",
            Router: `/testPaging`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/user`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/user/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePassword",
            Router: `/user/password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/BaseControllers:UserController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/user/resetPassword`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
