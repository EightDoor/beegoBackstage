package CommModels

import (
	"github.com/beego/beego/v2/core/validation"
	"time"
)

type BaseModel struct {
	Id        int       `json:"id" orm:"auto"`
	CreatedAt time.Time `json:"createdAt" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updatedAt" orm:"auto_now;type(datetime)"`
	DeletedAt time.Time `json:"deletedAt" orm:"-"`
}

// Valid 自定义校验
func (base *BaseModel) Valid(v *validation.Validation) {

}
