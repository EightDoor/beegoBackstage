package SysModels

import (
	"beegoBackstage/commModel"
)

type SysDict struct {
	commModel.BaseModel
	Name         string `json:"name" valid:"Required" label:"字典名称"`
	SerialNumber string `json:"serialNumber" valid:"Required" label:"字典编号"`
	Describe     string `json:"describe"`
}
