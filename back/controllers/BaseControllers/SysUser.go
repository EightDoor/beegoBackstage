package BaseControllers

import (
	"beegoBackstage/models/BaseModels"
	"beegoBackstage/utils"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type UserController struct {
	utils.BaseController
}

// Get
// @router /user [get]
func (c *UserController) Get() {
	var user []*BaseModels.SysUser
	o := orm.NewOrm()
	result := o.QueryTable(BaseModels.SysUser{})
	c.RPaging(utils.RPage{OrmResult: result, Data: &user})
}

// GetAll
// @router /user/all [get]
func (c *UserController) GetAll() {
	var user []*BaseModels.SysUser
	o := orm.NewOrm()
	_, err := o.QueryTable(BaseModels.SysUser{}).All(&user)
	c.RBack(utils.R{Data: user}, err)
}

// Post
// @router /user [post]
func (c *UserController) Post() {
	var user BaseModels.SysUser
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
	var user BaseModels.SysUser
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
	var user BaseModels.SysUser
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Error(err, "update-err")
	c.CustomValid(&user)
	if err != nil {
		c.RBack(utils.R{Data: "json format err"}, err)
	} else {
		msg, err := o.Update(&user, "account", "nickName", "email", "deptId", "phoneNum", "status", "avatar")
		c.RBack(utils.R{Data: msg}, err)
	}
}

// UpdatePassword 更新密码
// @router /user/password [post]
func (c *UserController) UpdatePassword() {
	var userPassword BaseModels.SysUserUpdatePassword
	var user BaseModels.SysUser
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &userPassword)
	// 表单校验
	c.CustomValid(userPassword)
	logs.Info(userPassword, "密码为")
	if err != nil {
		c.RBack(utils.R{Data: "json format err"}, err)
	} else {
		// 查询是否存在用户，密码是否正确
		qs := o.QueryTable(BaseModels.SysUser{})
		singPasswd := utils.PasswordEncryption(userPassword.Password)
		logs.Info(singPasswd, "加密的密码")
		userIsError := qs.Filter("id", userPassword.Id).Filter("password", singPasswd).One(&user)
		logs.Info(userIsError, "userIsError")
		if userIsError == nil && user.Account != "" {
			logs.Info(user, "是否查询到用户")
			if userPassword.NewPassword != "" {
				msg, err := o.QueryTable(BaseModels.SysUser{}).Filter("id", userPassword.Id).Update(orm.Params{
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
	var userPassword BaseModels.SysUserUpdatePassword
	// 读取body参数
	c.GetBodyToJson(&userPassword)
	// 表单验证
	c.CustomValid(&userPassword)
	o := orm.NewOrm()
	result, err := o.QueryTable(BaseModels.SysUser{}).Filter("id", userPassword.Id).Update(orm.Params{
		"password": utils.PasswordEncryption(userPassword.Password),
	})
	c.RBack(utils.R{Data: result}, err)
}
