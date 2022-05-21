package BaseModels

import CommModels "beegoBackstage/commModels"

type SysUser struct {
	CommModels.BaseModel
	Account  string `json:"account"`
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	DeptId   string `json:"deptId"`
	PhoneNum string `json:"phoneNum"`
	Status   int8   `json:"status"`
	Avatar   string `json:"avatar"`
	PassWord string `json:"passWord"`
}
