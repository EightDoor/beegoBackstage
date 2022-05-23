package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
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
			tokenStr, expirationTime := utils.GenerateToken(sysUser.Id)
			c.RSuccess(utils.R{
				Data: SysModels.LoginCall{Token: tokenStr, ExpirationTime: expirationTime},
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
	auth := c.Ctx.Request.Header.Get("Authorization")
	logs.Info(auth, "Authorization")
	if auth != "" {
		userId, err := utils.ValidateToken(auth)
		if err == nil {
			c.RSuccess(utils.R{Data: userId})
		} else {
			c.RError(utils.R{Msg: err.Error()})
		}
	} else {
		c.RError(utils.R{Msg: "Authorization不能为空"})
	}
}

// RefreshToken 刷新token
// @router /refreshToken [get]
func (c *LoginController) RefreshToken() {
	auth := c.Ctx.Request.Header.Get("Authorization")
	logs.Info(auth, "Authorization")
	if auth != "" {
		token, refreshTime, err := utils.RefreshToken(auth)
		if err == nil {
			loginCall := new(SysModels.LoginCall)
			loginCall.Token = token
			loginCall.ExpirationTime = time.Unix(refreshTime, 0)
			c.RSuccess(utils.R{Data: loginCall})
		} else {
			c.RError(utils.R{Msg: err.Error()})
		}
	} else {
		c.RError(utils.R{Msg: "Authorization不能为空"})
	}
}
