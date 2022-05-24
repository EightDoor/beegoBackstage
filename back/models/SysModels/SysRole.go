package SysModels

import (
	CommModels "beegoBackstage/commModels"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type SysRole struct {
	CommModels.BaseModel
	Remark   string `json:"remark"`
	RoleName string `json:"roleName" valid:"Required" label:"角色名称"`
}

// GetRoleList
// 根据用户id，查询角色列表
func GetRoleList(id int, list *[]SysRole) error {
	var userRole []SysUserRole
	var rError error
	o := orm.NewOrm()
	// 查询关联表菜单id 列表
	_, err := o.QueryTable(SysUserRole{}).Filter("userId", id).All(&userRole)
	if err == nil {
		var ids []int
		for _, v := range userRole {
			ids = append(ids, v.RoleId)
		}
		logs.Info(ids, "查询到的角色列表 - ids")
		if len(ids) > 0 {
			_, resultErr := o.QueryTable(SysRole{}).Filter("id__in", ids).All(list)
			if resultErr == nil {
				logs.Info(len(*list), "查询到的角色列表数量")
			} else {
				rError = resultErr
			}
		}
		rError = nil
	} else {
		rError = err
	}
	return rError
}
