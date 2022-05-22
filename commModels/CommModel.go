package CommModels

import "time"

type BaseModel struct {
	Id        int       `json:"id" orm:"auto"`
	CreatedAt time.Time `json:"createdAt" orm:"auto_now_add;"`
	UpdatedAt time.Time `json:"updatedAt" orm:"auto_now_update;type(datetime)"`
	DeletedAt time.Time `json:"deletedAt" orm:"auto_now_delete;type(datetime)"`
}
