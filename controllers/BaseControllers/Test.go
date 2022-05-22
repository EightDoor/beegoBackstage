package BaseControllers

import (
	"beegoBackstage/models/BaseModels"
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
	id, err := o.Insert(BaseModels.Test{Id: 1, Value: "5"})
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
	var t []*BaseModels.Test
	qs := o.QueryTable("test")
	qs.All(&t)
	c.RSuccess(utils.R{Data: t})
}

// GetPaging
// @router /testPaging [get]
func (c *TestController) GetPaging() {
	// orm查询数据
	o := orm.NewOrm()
	var dataInfo []*BaseModels.Test
	qs := o.QueryTable("test")
	c.RPaging(utils.RPage{
		OrmResult: qs,
		Data:      &dataInfo})
}
