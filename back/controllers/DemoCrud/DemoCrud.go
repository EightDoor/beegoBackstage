package DemoCrud

import (
	"beegoBackstage/models/DemoCrud"
	"beegoBackstage/utils"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type DemoCrudController struct {
	utils.BaseController
}

// Get
// @router /demoCrud [get]
func (c *DemoCrudController) Get() {
	var model []*DemoCrud.DemoCrud
	o := orm.NewOrm()
	result := o.QueryTable(DemoCrud.DemoCrud{})
	c.RPaging(utils.RPage{OrmResult: result, Data: &model})
}

// GetAll
// @router /demoCrud/all [get]
func (c *DemoCrudController) GetAll() {
	var model []*DemoCrud.DemoCrud
	o := orm.NewOrm()
	_, err := o.QueryTable(DemoCrud.DemoCrud{}).All(&model)
	c.RBack(utils.R{Data: model}, err)
}

// Post
// @router /demoCrud [post]
func (c *DemoCrudController) Post() {
	var model DemoCrud.DemoCrud
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
// @router /demoCrud/:id [delete]
func (c *DemoCrudController) Del() {
	var model DemoCrud.DemoCrud
	o := orm.NewOrm()
	// 读取参数id
	id, _ := c.GetInt(":id")
	model.Id = id
	logs.Info(model.Id, "id")
	result, err := o.Delete(&model)
	c.RBack(utils.R{Data: result}, err)
}

// Update
// @router /demoCrud [put]
func (c *DemoCrudController) Update() {
	var model DemoCrud.DemoCrud
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	logs.Error(err, "update-err")
	c.CustomValid(&model)
	if err != nil {
		c.RBack(utils.R{Data: "json format err"}, err)
	} else {
		msg, err := o.Update(&model)
		c.RBack(utils.R{Data: msg}, err)
	}
}
