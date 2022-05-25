package models

import (
	"beegoBackstage/models/DemoCrud"
	"beegoBackstage/models/JournalModels"
	"beegoBackstage/models/SysModels"
	"github.com/beego/beego/v2/client/orm"
)

// RegisterModels orm使用，需要注册Model
func RegisterModels() {
	orm.RegisterModel(
		// sys
		new(SysModels.Test),
		new(SysModels.SysUser),
		new(SysModels.SysRole),
		new(SysModels.SysMenu),
		new(SysModels.SysDept),
		new(SysModels.SysDict),
		new(SysModels.SysDictItem),
		new(SysModels.SysUserRole),
		new(SysModels.SysRoleMenu),
		// user log
		new(JournalModels.LogLogin),

		// other
		new(DemoCrud.DemoCrud),
	)
}
