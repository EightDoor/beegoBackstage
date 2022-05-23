package SysModels

import CommModels "beegoBackstage/commModels"

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
