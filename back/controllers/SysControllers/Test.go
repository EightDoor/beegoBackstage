package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/client/orm"
)

type TestController struct {
	utils.BaseController
}

// Post
// @router /test [post]
func (c *TestController) Post() {
	// orm查询数据
	o := orm.NewOrm()
	id, err := o.Insert(SysModels.Test{Id: 1, Value: "5"})
	if err != nil {
		c.Data["json"] = id
	}
	c.ServeJSON()
}

// Get
// @router /test [get]
func (c *TestController) Get() {
	// orm查询数据
	o := orm.NewOrm()
	var t []*SysModels.Test
	qs := o.QueryTable("test")
	qs.All(&t)
	c.RSuccess(utils.R{Data: t})
}

// GetPaging
// @router /testPaging [get]
func (c *TestController) GetPaging() {
	// orm查询数据
	o := orm.NewOrm()
	var dataInfo []*SysModels.Test
	qs := o.QueryTable("test")
	c.RPaging(utils.RPage{
		OrmResult: qs,
		Data:      &dataInfo})
}
