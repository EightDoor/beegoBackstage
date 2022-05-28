// 状态码从100开始
export enum RequestStatusMsg {
  UNKNOWN_ERROR = 100, // 未知错误
  RequestAuthorizedFailed = 101, // 登录token鉴权失败
}
export const TOKEN = 'TOKEN'
// 选中的面包屑列表
export const PERMISSIONBUTTONS = 'PERMISSIONBUTTONS'

// 当前选择的菜单项
export const CURRENT_MENU = 'CURRENT_MENU'
// 存储的所有菜单项列表
export const LIST_OF_ALL_STORED_MENU_ITEMS = 'LIST_OF_ALL_STORED_MENU_ITEMS'
