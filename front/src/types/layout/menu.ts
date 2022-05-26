export interface MenuItem {
  key: number
  title: string
  icon?: string
  path?: string
  children?: MenuItem[]
  id: number
  parentId: number
  crumb?: string
  closable?: boolean
  isHome?: number
  name?: string
}
export interface MenusInfo {
  selectedKeys: number[]
  list: MenuItem[]
  openKeys: number[]
}
