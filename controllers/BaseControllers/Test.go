package BaseControllers

import (
	"beegoBackstage/models/BaseModels"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type TestController struct {
	web.Controller
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
	t := new([]BaseModels.Test)
	err := o.Read(&t)
	if err != nil {
		c.Data["json"] = t
	}
	c.Data["json"] = "出错了"
	c.ServeJSON()
}
