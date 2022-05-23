package SysModels

import CommModels "beegoBackstage/commModels"

type SysUser struct {
	CommModels.BaseModel
	Account  string `json:"account" valid:"Required" label:"账户名称"`
	NickName string `json:"nickName" valid:"Required" label:"姓名"`
	Email    string `json:"email"`
	DeptId   string `json:"deptId" valid:"Required" label:"部门id"`
	PhoneNum string `json:"phoneNum"`
	Status   int8   `json:"status"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

type SysUserUpdatePassword struct {
	Id          int    `json:"id" valid:"Required" label:"用户主键id"`
	Password    string `json:"password" valid:"Required" label:"旧密码"`
	NewPassword string `json:"newPassword"`
}
