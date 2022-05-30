package commModel

import (
	"github.com/beego/beego/v2/core/validation"
	"time"
)

type BaseModel struct {
	Id        int        `json:"id" orm:"auto"`
	CreatedAt *time.Time `json:"createdAt,omitempty" orm:"auto_now_add;type(datetime)"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" orm:"auto_now;type(datetime);"`
	DeletedAt *time.Time `json:"-" orm:"-"`
}

// Valid 自定义校验
func (base *BaseModel) Valid(v *validation.Validation) {

}
