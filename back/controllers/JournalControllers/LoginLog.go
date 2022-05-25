package JournalControllers

import (
	"beegoBackstage/models/JournalModels"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type LoginLogController struct {
	utils.BaseController
}

// Get 查询列表
// @router /loginLog [get]
func (c *LoginLogController) Get() {
	var model []JournalModels.LogLogin
	o := orm.NewOrm()
	qs := o.QueryTable(JournalModels.LogLogin{}).RelatedSel()
	logs.Info(model, "model")
	c.RPaging(utils.RPage{
		OrmResult: qs,
		Data:      &model,
	})
}
