package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type DictItemController struct {
	utils.BaseController
}

// Get
// @router /dictItem [get]
func (c *DictItemController) Get() {
	var model []*SysModels.SysDictItem
	o := orm.NewOrm()
	result := o.QueryTable(SysModels.SysDictItem{})
	c.RPaging(utils.RPage{OrmResult: result, Data: &model})
}

// GetAll
// @router /dictItem/all [get]
func (c *DictItemController) GetAll() {
	var model []*SysModels.SysDictItem
	o := orm.NewOrm()
	result, err := o.QueryTable(SysModels.SysDictItem{}).All(&model)
	c.RBack(utils.R{Data: result}, err)
}

// Post
// @router /dictItem [post]
func (c *DictItemController) Post() {
	var model SysModels.SysDictItem
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
// @router /dictItem/:id [delete]
func (c *DictItemController) Del() {
	var model SysModels.SysDictItem
	o := orm.NewOrm()
	// 读取参数id
	id, _ := c.GetInt(":id")
	model.Id = id
	logs.Info(model.Id, "id")
	result, err := o.Delete(&model)
	c.RBack(utils.R{Data: result}, err)
}

// Update
// @router /dictItem [put]
func (c *DictItemController) Update() {
	var user SysModels.SysDictItem
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
