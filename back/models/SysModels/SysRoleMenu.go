package SysModels

import (
	"beegoBackstage/commModel"
)

type SysRoleMenu struct {
	commModel.BaseModel
	RoleId int `json:"roleId" valid:"Required" label:"角色id"`
	MenuId int `json:"menuId" valid:"Required" label:"菜单id"`
}

type SysRoleMenuGet struct {
	RoleId int   `json:"roleId" valid:"Required" label:"角色id"`
	MenuId []int `json:"menuId" valid:"Required" label:"菜单id"`
}
