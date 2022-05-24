package SysModels

import CommModels "beegoBackstage/commModels"

type SysDictItem struct {
	CommModels.BaseModel
	Label    string `json:"label" valid:"Required" label:"名称"`
	Value    string `json:"value" valid:"Required" label:"数据值"`
	DictId   int    `json:"dictId" valid:"Required" label:"字典id"`
	Describe string `json:"describe"`
}
