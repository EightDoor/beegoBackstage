package SysModels

import CommModels "beegoBackstage/commModels"

type SysDept struct {
	CommModels.BaseModel
	OrderNum int    `json:"orderNum"`
	Name     string `json:"name" valid:"Required" label:"部门名称"`
}
