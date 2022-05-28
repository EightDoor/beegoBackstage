import { BREAD_CRUMBS, OPEN_LEFT_MENU } from '@/store/mutation-types'
import type { MenuType } from '@/types/sys'

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
  menuOpenKeys: string[]
}

export default {
  namespace: true,
  state: {
    panes: [],
    list: [],
    menuOpenKeys: [],
  } as CrumbsStoreType,
  mutations: {
    [BREAD_CRUMBS](state: CrumbsStoreType, payload: string[]) {
      state.list = payload
    },
    [OPEN_LEFT_MENU](state: CrumbsStoreType, payload: MenuType) {
      if (payload.menuOpenKeys)
        state.menuOpenKeys = payload.menuOpenKeys.split(',')
      else
        state.menuOpenKeys = []
    },
  },
}
