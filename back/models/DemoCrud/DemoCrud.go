package DemoCrud

import CommModels "beegoBackstage/commModels"

type DemoCrud struct {
	CommModels.BaseModel
	Title   string `json:"title" valid:"Required" label:"标题"`
	Content string `json:"content" valid:"Required" label:"内容"`
	Type    string `json:"type" valid:"Required" label:"类型"`
}
