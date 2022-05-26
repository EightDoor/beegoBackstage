export interface PanesType {
  id: number
  title: string
  path: string
  parentId: number
  closable?: boolean
  timestamp?: number
}
export interface CrumbsStoreType {
  // tabs
  panes: PanesType[]
  // 面包屑
  list: string[]
  // 菜单列表展开 keys
  menuOpenKeys: number[]
}

export default {
  namespace: true,
  state: {
    panes: [],
    list: [],
    menuOpenKeys: [],
  } as CrumbsStoreType,
}
