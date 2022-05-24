package SysModels

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type SysUserInfo struct {
	UserInfo SysUser   `json:"userInfo"`
	Menus    []SysMenu `json:"menus"`
	Roles    []SysRole `json:"roles"`
}

// GetUserInfoData
// 根据用户id获取用户信息
func GetUserInfoData(id int) (SysUserInfo, error) {
	var sysUserInfo SysUserInfo
	var sysMenu []SysMenu
	var sysRole []SysRole
	var sysRoleMenu []SysRoleMenu
	var sysUser SysUser
	o := orm.NewOrm()
	// 查询用户信息
	userErr := o.QueryTable(SysUser{}).Filter("id", id).One(&sysUser)
	if userErr != nil {
		return sysUserInfo, userErr
	}
	sysUserInfo.UserInfo = sysUser
	// 查询用户角色关联表
	var sysUserRole []SysUserRole
	_, err := o.QueryTable(SysUserRole{}).Filter("userId", id).All(&sysUserRole)
	logs.Info(sysUserRole, "用户角色")
	if err != nil {
		return sysUserInfo, err
	}
	// 角色id列表
	var ids []int
	// 查询用户角色
	for _, v := range sysUserRole {
		ids = append(ids, v.RoleId)
	}
	_, roleErr := o.QueryTable(SysRole{}).Filter("id__in", ids).All(&sysRole)
	logs.Info(sysRole, "角色列表")
	sysUserInfo.Roles = sysRole
	if roleErr != nil {
		return sysUserInfo, err
	}
	// 根据角色列表，查询用户菜单列表
	_, roleMenuERr := o.QueryTable(SysRoleMenu{}).Filter("roleId__in", ids).All(&sysRoleMenu)
	if roleMenuERr != nil {
		return sysUserInfo, roleMenuERr
	}
	// 菜单列表ids
	var roleMenuIds []int
	for _, roleVal := range sysRoleMenu {
		roleMenuIds = append(roleMenuIds, roleVal.MenuId)
	}
	// 查询菜单列表
	_, menusErr := o.QueryTable(SysMenu{}).Filter("id__in", roleMenuIds).All(&sysMenu)
	if menusErr != nil {
		return sysUserInfo, menusErr
	}
	sysUserInfo.Menus = sysMenu
	logs.Info(sysMenu, "用户拥有菜单、角色列表")
	return sysUserInfo, nil
}
