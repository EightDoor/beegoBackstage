import type { CommonTableList, CommonTreeSelect } from '@/types/type'

export interface MenuType extends CommonTableList, ITabs {
  crumbs?: string
  key?: number
  id: number
  createdAt?: any
  updatedAt?: any
  deletedAt?: any
  parentId: number
  title: string
  type?: number
  orderNum?: number
  perms?: any
  name?: string | number
  path?: string | number
  component?: string
  redirect?: any
  icon?: any
  isHome?: number
  children?: Array<MenuType>
  value?: string
  hidden?: number
  menuId?: string
}

export interface DepartType extends CommonTreeSelect {
  parentId: string
  name: string
  orderNum: number
  id?: string
  children?: Array<MenuType>
}

export interface UserType extends CommonTableList {
  deletedAt?: any
  account: string
  nickName: string
  email?: any
  status: number
  avatar?: any
  deptId: number
  phoneNum?: any
  password?: string
}

export interface RoleType extends CommonTableList {
  // 描述
  remark?: string
  // 角色名称
  roleName: string
  id?: number
}

export interface LoginType {
  account: string
  pass_word: string
}

interface Role {
  id: number
  createdAt?: any
  updatedAt?: any
  deletedAt?: any
  remark: string
  roleName: string
}

export interface UserInformation {
  userInfo: UserType
  menus: MenuType[]
  roles: Role[]
}

export interface ITabs {
  closable?: boolean
}
