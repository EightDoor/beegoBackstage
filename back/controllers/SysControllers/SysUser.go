package SysControllers

import (
	"beegoBackstage/models"
	"beegoBackstage/models/JournalModels"
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	utils.BaseController
}

// Get
// @router /user [get]
func (c *UserController) Get() {
	var user []*SysModels.SysUser
	o := orm.NewOrm()
	result := o.QueryTable(SysModels.SysUser{}).RelatedSel()
	c.RPaging(utils.RPage{OrmResult: result, Data: &user})
}

// GetAll
// @router /user/all [get]
func (c *UserController) GetAll() {
	var user []*SysModels.SysUser
	o := orm.NewOrm()
	_, err := o.QueryTable(SysModels.SysUser{}).All(&user)
	c.RBack(utils.R{Data: user}, err)
}

// Post
// @router /user [post]
func (c *UserController) Post() {
	var user SysModels.SysUser
	o := orm.NewOrm()
	// 解析body参数
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	c.CustomValid(&user)
	user.Password = utils.PasswordEncryption(user.Password)
	result, err := o.Insert(&user)
	if err == nil && result > 0 {
		c.RSuccess(utils.R{Data: user})
	} else {
		c.RError(utils.R{Data: err.Error()})
	}
}

// Del
// @router /user/:id [delete]
func (c *UserController) Del() {
	var user SysModels.SysUser
	o := orm.NewOrm()
	// 读取参数id
	id, _ := c.GetInt(":id")
	user.Id = id
	logs.Info(user.Id, "id")
	result, err := o.Delete(&user)
	c.RBack(utils.R{Data: result}, err)
}

// Update
// @router /user [put]
func (c *UserController) Update() {
	var user SysModels.SysUser
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Error(err, "update-err")
	c.CustomValid(&user)
	if err != nil {
		c.RBack(utils.R{Data: "json format err"}, err)
	} else {
		msg, err := o.Update(&user, "account", "nickName", "email", "deptId", "phoneNum", "status", "file")
		c.RBack(utils.R{Data: msg}, err)
	}
}

// UpdatePassword 更新密码
// @router /user/password [post]
func (c *UserController) UpdatePassword() {
	var userPassword SysModels.SysUserUpdatePassword
	var user SysModels.SysUser
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &userPassword)
	// 表单校验
	c.CustomValid(userPassword)
	logs.Info(userPassword, "密码为")
	if err != nil {
		c.RBack(utils.R{Data: "json format err"}, err)
	} else {
		// 查询是否存在用户，密码是否正确
		qs := o.QueryTable(SysModels.SysUser{})
		singPasswd := utils.PasswordEncryption(userPassword.Password)
		logs.Info(singPasswd, "加密的密码")
		userIsError := qs.Filter("id", userPassword.Id).Filter("password", singPasswd).One(&user)
		logs.Info(userIsError, "userIsError")
		if userIsError == nil && user.Account != "" {
			logs.Info(user, "是否查询到用户")
			if userPassword.NewPassword != "" {
				msg, err := o.QueryTable(SysModels.SysUser{}).Filter("id", userPassword.Id).Update(orm.Params{
					"password": utils.PasswordEncryption(userPassword.NewPassword),
				})
				c.RBack(utils.R{Data: msg}, err)
			} else {
				c.RError(utils.R{Msg: "请填写新密码"})
			}
		} else {
			c.RBack(utils.R{Data: userIsError}, errors.New("查询用户密码 失败"))
		}
	}
}

// ResetPassword
// @router /user/resetPassword [post]
func (c *UserController) ResetPassword() {
	var userPassword SysModels.SysUserUpdatePassword
	// 读取body参数
	c.GetBodyToJson(&userPassword)
	// 表单验证
	c.CustomValid(&userPassword)
	o := orm.NewOrm()
	result, err := o.QueryTable(SysModels.SysUser{}).Filter("id", userPassword.Id).Update(orm.Params{
		"password": utils.PasswordEncryption(userPassword.Password),
	})
	c.RBack(utils.R{Data: result}, err)
}

// UserRoleList 查询用户关联角色列表
// @router /userRoleList/:id [get]
func (c *UserController) UserRoleList() {
	var role []SysModels.SysRole
	id, err := c.GetInt(":id")
	if err == nil {
		err := SysModels.GetRoleList(id, &role)
		if err == nil {
			c.RSuccess(utils.R{
				Data: role,
			})
		} else {
			c.RError(utils.R{
				Msg: err.Error(),
			})
		}
	} else {
		c.RError(utils.R{
			Data: "请传递查询路劲id, userRoleList/:id",
		})
	}
}

// GetTokenUserInfo 根据token获取用户菜单、角色
// @router /userInfo [get]
func (c *UserController) GetTokenUserInfo() {
	token := c.Ctx.Input.Header("Authorization")
	if token != "" {
		userId, userIdErr := utils.ValidateToken(token)

		// 用户log日志
		userLog, _ := web.AppConfig.Bool("userLog")
		logs.Info(userLog, "userLog")
		if userLog {
			var logLogin JournalModels.LogLogin
			var sysUser SysModels.SysUser
			sysUser.Id = userId
			logLogin.User = &sysUser
			logLogin.Ip = c.Ctx.Input.Header("customIp")
			logLogin.Equipment = c.Ctx.Input.Header("customOs")
			JournalModels.LogLoginCreate(logLogin)
		}
		if userIdErr == nil {
			result, sysUserInfoDataErr := SysModels.GetUserInfoData(userId)
			if sysUserInfoDataErr == nil {
				c.RSuccess(utils.R{Data: result})
			} else {
				c.RError(utils.R{Msg: sysUserInfoDataErr.Error(), Data: "查询出错"})
			}
		} else {
			c.RError(utils.R{
				Msg:  userIdErr.Error(),
				Code: models.NO_AUTHORIZATION,
			})
		}
	} else {
		c.RError(utils.R{
			Msg:  "headers Authorization 不存在",
			Code: models.NO_AUTHORIZATION,
		})
	}
}
