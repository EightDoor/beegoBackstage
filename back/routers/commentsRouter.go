package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/dept`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/dept`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/dept`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/dept/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DeptController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/dept/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/dict`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/dict`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/dict`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/dict/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/dict/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictController"],
        beego.ControllerComments{
            Method: "DictQueryItem",
            Router: `/dictQueryItem/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/dictItem`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/dictItem`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/dictItem`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/dictItem/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:DictItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/dictItem/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:LoginController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:LoginController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/menu`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/menu`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/menu`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/menu/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:MenuController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/menu/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/role`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/role`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/role`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/role/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/role/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:RoleController"],
        beego.ControllerComments{
            Method: "RoleMenuRelationList",
            Router: `/roleMenuRelationList/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:SysRoleMenuController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:SysRoleMenuController"],
        beego.ControllerComments{
            Method: "RoleMenuRelation",
            Router: `/roleMenuRelation`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:SysUserRoleController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:SysUserRoleController"],
        beego.ControllerComments{
            Method: "Relation",
            Router: `/userRoleRelation`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:TestController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/test`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:TestController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:TestController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:TestController"],
        beego.ControllerComments{
            Method: "GetPaging",
            Router: `/testPaging`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/user`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/user/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePassword",
            Router: `/user/password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/user/resetPassword`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"] = append(beego.GlobalControllerRouter["beegoBackstage/controllers/SysControllers:UserController"],
        beego.ControllerComments{
            Method: "UserRoleList",
            Router: `/userRoleList/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
