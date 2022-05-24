package SysModels

import (
	CommModels "beegoBackstage/commModels"
	"github.com/beego/beego/v2/client/orm"
)

type SysMenu struct {
	CommModels.BaseModel
	OrderNum int    `json:"orderNum"`
	Redirect string `json:"redirect"`
	Icon     string `json:"icon"`
	Hidden   int    `json:"hidden"`
	IsHome   int    `json:"isHome"`
	ParentId int    `json:"parentId"`
	Title    string `json:"title" valid:"Required" label:"标题"`
	Type     int    `json:"type" valid:"Required" label:"菜单类型： 1. 目录 2. 菜单  3. 按钮"`
	Perms    string `json:"perms"`
	Name     string `json:"name"`
}

// GetRoleIdMenus
// 根据角色id 查询菜单列表
func GetRoleIdMenus(id int, list *[]SysMenu) error {
	var rError error
	var sysRoleMenu []SysRoleMenu
	// 查询关联表ids
	var ids []int
	o := orm.NewOrm()
	_, err := o.QueryTable(SysRoleMenu{}).Filter("roleId", id).All(&sysRoleMenu)
	if err == nil {
		for _, v := range sysRoleMenu {
			ids = append(ids, v.MenuId)
		}
		if len(ids) > 0 {
			_, resultError := o.QueryTable(SysMenu{}).Filter("id__in", ids).All(list)
			rError = resultError
		} else {
			rError = nil
		}

	} else {
		rError = err
	}
	return rError
}
