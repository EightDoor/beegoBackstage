package JournalModels

import (
	"beegoBackstage/commModel"
	"beegoBackstage/models/SysModels"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type LogLogin struct {
	commModel.BaseModel
	Equipment string             `json:"equipment"`
	User      *SysModels.SysUser `json:"user" orm:"rel(fk)"`
	Ip        string             `json:"ip"`
}

// LogLoginCreate
// 创建记录
func LogLoginCreate(data LogLogin) (int64, error) {
	o := orm.NewOrm()
	result, err := o.Insert(&data)
	if err != nil {
		logs.Error(err, "创建日志记录")
	}
	return result, err
}
