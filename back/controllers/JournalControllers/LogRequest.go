package JournalControllers

import (
	"beegoBackstage/models/JournalModels"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/client/orm"
)

type LogRequestController struct {
	utils.BaseController
}

// Get 请求日志
// @router /logRequest [get]
func (c *LogRequestController) Get() {
	var logRequest []JournalModels.LogRequest
	o := orm.NewOrm()
	qs := o.QueryTable(JournalModels.LogRequest{}).RelatedSel()
	c.RPaging(utils.RPage{
		OrmResult: qs,
		Data:      &logRequest,
	})
}
