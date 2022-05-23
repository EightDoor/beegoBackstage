package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type RoleController struct {
	utils.BaseController
}

// Get
// @router /role [get]
func (c *RoleController) Get() {
	var model []*SysModels.SysRole
	o := orm.NewOrm()
	result := o.QueryTable(SysModels.SysRole{})
	c.RPaging(utils.RPage{OrmResult: result, Data: &model})
}

// GetAll
// @router /role/all [get]
func (c *RoleController) GetAll() {
	var model []*SysModels.SysRole
	o := orm.NewOrm()
	result, err := o.QueryTable(SysModels.SysRole{}).All(&model)
	c.RBack(utils.R{Data: result}, err)
}

// Post
// @router /role [post]
func (c *RoleController) Post() {
	var model SysModels.SysRole
	o := orm.NewOrm()
	// 解析body参数
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	c.CustomValid(&model)
	result, err := o.Insert(&model)
	if err == nil && result > 0 {
		c.RSuccess(utils.R{Data: model})
	} else {
		c.RError(utils.R{Data: err.Error()})
	}
}

// Del
// @router /role/:id [delete]
func (c *RoleController) Del() {
	var model SysModels.SysRole
	o := orm.NewOrm()
	// 读取参数id
	id, _ := c.GetInt(":id")
	model.Id = id
	logs.Info(model.Id, "id")
	result, err := o.Delete(&model)
	c.RBack(utils.R{Data: result}, err)
}

// Update
// @router /role [put]
func (c *RoleController) Update() {
	var user SysModels.SysRole
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Error(err, "update-err")
	c.CustomValid(&user)
	if err != nil {
		c.RBack(utils.R{Data: "json format err"}, err)
	} else {
		msg, err := o.Update(&user)
		c.RBack(utils.R{Data: msg}, err)
	}
}

// RoleMenuRelationList 角色关联菜单列表
// @router /roleMenuRelationList/:id [get]
func (c *RoleController) RoleMenuRelationList() {
	var sysMenu []SysModels.SysMenu
	id, err := c.GetInt(":id")
	if err == nil {
		rError := SysModels.GetRoleIdMenus(id, &sysMenu)
		if rError == nil {
			c.RSuccess(utils.R{
				Data: sysMenu,
			})
		} else {
			c.RError(utils.R{
				Msg: rError.Error(),
			})
		}
	} else {
		c.RError(utils.R{
			Msg: err.Error(),
		})
	}
}
