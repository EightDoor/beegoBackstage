package DemoCrud

import (
	"beegoBackstage/commModel"
)

type DemoCrud struct {
	commModel.BaseModel
	Title   string `json:"title" valid:"Required" label:"标题"`
	Content string `json:"content" valid:"Required" label:"内容"`
	Type    string `json:"type" valid:"Required" label:"类型"`
}
