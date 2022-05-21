package BaseModels

import (
	"github.com/beego/beego/v2/client/orm"
)

func RegisterModels() {
	orm.RegisterModel(new(Test))
}
