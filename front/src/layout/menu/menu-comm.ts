import { toRaw } from 'vue'
import { CURRENT_MENU } from '@/utils/constant'
import storeInstant from '@/utils/store'
import type { MenuItem } from '@/types/layout/menu'

export const commSetData = async (item: MenuItem) => {
  // 当前选中菜单,用于处理当前菜单的按钮权限
  await storeInstant.set(CURRENT_MENU, toRaw(item))
}
