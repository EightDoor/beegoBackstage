package SysModels

import CommModels "beegoBackstage/commModels"

type SysUserRole struct {
	CommModels.BaseModel
	UserId int `json:"userId" valid:"Required" label:"用户id"`
	RoleId int `json:"roleId" valid:"Required" label:"角色id"`
}
type SysUserRoleGet struct {
	UserId int   `json:"userId" valid:"Required" label:"用户id"`
	RoleId []int `json:"roleId" valid:"Required" label:"角色id"`
}
