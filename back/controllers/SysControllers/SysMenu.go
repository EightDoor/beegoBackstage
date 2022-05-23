package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type MenuController struct {
	utils.BaseController
}

// Get
// @router /menu [get]
func (c *MenuController) Get() {
	var model []*SysModels.SysMenu
	o := orm.NewOrm()
	result := o.QueryTable(SysModels.SysMenu{})
	c.RPaging(utils.RPage{OrmResult: result, Data: &model})
}

// GetAll
// @router /menu/all [get]
func (c *MenuController) GetAll() {
	var model []*SysModels.SysMenu
	o := orm.NewOrm()
	result, err := o.QueryTable(SysModels.SysMenu{}).All(&model)
	c.RBack(utils.R{Data: result}, err)
}

// Post
// @router /menu [post]
func (c *MenuController) Post() {
	var model SysModels.SysMenu
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
// @router /menu/:id [delete]
func (c *MenuController) Del() {
	var model SysModels.SysMenu
	o := orm.NewOrm()
	// 读取参数id
	id, _ := c.GetInt(":id")
	model.Id = id
	logs.Info(model.Id, "id")
	result, err := o.Delete(&model)
	c.RBack(utils.R{Data: result}, err)
}

// Update
// @router /menu [put]
func (c *MenuController) Update() {
	var user SysModels.SysMenu
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
