import { toRaw } from 'vue'
import { CURRENT_MENU } from '@/utils/constant'
import storeInstant from '@/utils/store'
import type { MenuItem } from '@/types/layout/menu'
import store from '@/store'
import { BREAD_CRUMBS } from '@/store/mutation-types'
import log from '@/utils/log'

export const commSetData = async (item: MenuItem) => {
  log.d(item, '点击菜单')
  // 当前选中菜单,用于处理当前菜单的按钮权限
  if (item.crumb) {
    const list = item.crumb.split(',')
    store.commit(BREAD_CRUMBS, list)
  }

  await storeInstant.set(CURRENT_MENU, toRaw(item))
}
