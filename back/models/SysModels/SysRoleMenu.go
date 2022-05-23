package SysModels

import CommModels "beegoBackstage/commModels"

type SysRoleMenu struct {
	CommModels.BaseModel
	RoleId int `json:"roleId" valid:"Required" label:"角色id"`
	MenuId int `json:"menuId" valid:"Required" label:"菜单id"`
}

type SysRoleMenuGet struct {
	RoleId int   `json:"roleId" valid:"Required" label:"角色id"`
	MenuId []int `json:"menuId" valid:"Required" label:"菜单id"`
}
