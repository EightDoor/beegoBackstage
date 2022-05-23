package SysControllers

import (
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type SysUserRoleController struct {
	utils.BaseController
}

// Relation 用户关联角色
// @router /userRoleRelation [post]
func (c *SysUserRoleController) Relation() {
	var userRoleGet SysModels.SysUserRoleGet
	c.GetBodyToJson(&userRoleGet)
	c.CustomValid(&userRoleGet)
	o := orm.NewOrm()
	// 使用事务，先删除用户全部相关的，然后再次添加
	to, err := o.Begin()
	if err != nil {
		logs.Error("开启事务失败")
		c.RError(utils.R{Msg: "开启事务失败"})
	}
	_, delError := o.QueryTable(SysModels.SysUserRole{}).Filter("userId", userRoleGet.UserId).Delete()
	if delError == nil {
		var userRole []SysModels.SysUserRole
		for _, v := range userRoleGet.RoleId {
			userRole = append(userRole, SysModels.SysUserRole{UserId: userRoleGet.UserId, RoleId: v})
		}
		// 新增传递的项
		result, insertErr := o.InsertMulti(len(userRole), &userRole)
		if insertErr == nil {
			to.Commit()
			c.RSuccess(utils.R{Data: result})
		} else {
			c.RError(utils.R{Msg: insertErr.Error()})
		}
	} else {
		logs.Error(delError, "执行失败")
		to.Rollback()
	}
}
