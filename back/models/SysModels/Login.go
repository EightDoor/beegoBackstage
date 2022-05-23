package SysModels

type Login struct {
	UserName string `json:"userName" valid:"Required" label:"用户名"`
	Password string `json:"password" valid:"Required" label:"密码"`
}

type LoginCall struct {
	Token string `json:"token"`
}
