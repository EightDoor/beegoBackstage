package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type SysRoleMenuController struct {
	utils.BaseController
}

// RoleMenuRelation 角色关联菜单
// @router /roleMenuRelation [post]
func (c *SysRoleMenuController) RoleMenuRelation() {
	var sysRoleMenuGet SysModels.SysRoleMenuGet
	c.GetBodyToJson(&sysRoleMenuGet)
	c.CustomValid(&sysRoleMenuGet)

	o := orm.NewOrm()
	// 使用事务，先删除角色相关联的菜单，然后添加
	to, err := o.Begin()

	if err == nil {
		// 删除关联
		_, delErr := o.QueryTable(SysModels.SysRoleMenu{}).Filter("roleId", sysRoleMenuGet.RoleId).Delete()
		if delErr == nil {
			// 添加新项
			var sysRoleMenu []SysModels.SysRoleMenu
			for _, v := range sysRoleMenuGet.MenuId {
				sysRoleMenu = append(sysRoleMenu, SysModels.SysRoleMenu{
					RoleId: sysRoleMenuGet.RoleId,
					MenuId: v,
				})
			}
			result, insertErr := o.InsertMulti(len(sysRoleMenu), &sysRoleMenu)
			if insertErr == nil {
				c.RSuccess(utils.R{Data: result})
			} else {
				logs.Error(insertErr, "添加失败")
				to.Rollback()
			}
		} else {
			to.Rollback()
			logs.Error(delErr, "删除失败,回滚")
		}
	} else {
		logs.Error("启用事务失败")
	}
}
