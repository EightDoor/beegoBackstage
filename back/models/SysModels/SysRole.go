package SysModels

import CommModels "beegoBackstage/commModels"

type SysRole struct {
	CommModels.BaseModel
	Remark   string `json:"remark"`
	RoleName string `json:"roleName" valid:"Required" label:"角色名称"`
}
