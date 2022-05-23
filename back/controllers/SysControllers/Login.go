package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/client/orm"
)

type LoginController struct {
	utils.BaseController
}

// Get
// @router / [post]
func (c *LoginController) Get() {
	var loginModel SysModels.Login
	var sysUser SysModels.SysUser
	c.GetBodyToJson(&loginModel)
	c.CustomValid(&loginModel)

	o := orm.NewOrm()
	isUserErr := o.QueryTable(SysModels.SysUser{}).Filter("account", loginModel.UserName).Filter("password", utils.PasswordEncryption(loginModel.Password)).One(&sysUser)
	if isUserErr == nil {
		if sysUser.Account != "" {
			// 登录 返回token
			c.RSuccess(utils.R{
				Data: SysModels.LoginCall{Token: utils.GenerateToken(sysUser.Id)},
			})
		} else {
			c.RError(utils.R{Msg: "用户名或者密码错误"})
		}
	} else {
		c.RError(utils.R{Msg: "用户名或者密码错误"})
	}
}

// ValidateToken 验证token是否有效
// @router /validateToken [get]
func (c *LoginController) ValidateToken() {
	
}
