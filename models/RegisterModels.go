package models

import (
	"beegoBackstage/models/BaseModels"
	"github.com/beego/beego/v2/client/orm"
)

// RegisterModels orm使用，需要注册Model
func RegisterModels() {
	orm.RegisterModel(new(BaseModels.Test), new(BaseModels.SysUser))
}
