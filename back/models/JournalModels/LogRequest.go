package JournalModels

import (
	CommModels "beegoBackstage/commModels"
	"beegoBackstage/models/SysModels"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type LogRequest struct {
	CommModels.BaseModel
	Params         string             `json:"params"`
	RequestAddress string             `json:"requestAddress"`
	Ip             string             `json:"ip"`
	User           *SysModels.SysUser `json:"user" orm:"rel(fk)"`
	Method         string             `json:"method"`
}

// LogRequestCreate
// 添加请求记录
func LogRequestCreate(logRequest LogRequest) {
	o := orm.NewOrm()
	_, err := o.Insert(&logRequest)
	if err != nil {
		logs.Error(err, "新增请求记录失败")
	}
}
