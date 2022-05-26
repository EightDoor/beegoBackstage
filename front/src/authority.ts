import type { RouteLocationNormalized, RouteRecordRaw, Router } from 'vue-router'
import { SET_SYS } from '@/store/mutation-types'
import store from '@/store/index'
import { TOKEN } from '@/utils/constant'
import log from '@/utils/log'

// 路由白名单
const routerWhiteList = ['/login']
export const canUserAccess = async (
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  router: Router,
): Promise<boolean> => {
  const status = routerWhiteList.find(item => to.path === item)
  const token = localStorage.getItem(TOKEN)
  if (!status) {
    if (token) {
      if (store.state.sys.userInfo && store.state.sys.userInfo.account)
        return true
      try {
        // 判断vuex是否存在值
        const { menus } = await store.dispatch(SET_SYS)
        log.i(menus, 'menus')
        menus.forEach((item: RouteRecordRaw) => {
          router.addRoute(item)
        })
        await router.push({ path: to.path, replace: true })
        return true
      }
      catch (err) {
        return to.name === 'login'
      }
    }
    else {
      await router.replace('/login')
      return true
    }
  }
  else {
    return true
  }
}
