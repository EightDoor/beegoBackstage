import type { Commit } from 'vuex'
import type { RouteRecordRaw } from 'vue-router'
import type { Key } from 'ant-design-vue/es/_util/type'
import { message } from 'ant-design-vue'
import { cloneDeep } from 'lodash-es'

import store from '../index'
import {
  BREAD_CRUMBS,
  COLLAPSED,
  LOGIN,
  OPEN_LEFT_MENU,
  SETUSERINFO,
  SET_MENUS_MUTATION, SET_SYS, USERINFOMENUS,
} from '@/store/mutation-types'
import type { LoginType, MenuType, UserInformation, UserType } from '@/types/sys'
import http from '@/utils/request'
import {
  CURRENT_MENU,
  LIST_OF_ALL_STORED_MENU_ITEMS,
  PERMISSIONBUTTONS,
  TOKEN,
} from '@/utils/constant'

import { ListObjCompare, ListToTree } from '@/utils'
import log from '@/utils/log'
import storeInstant from '@/utils/store'

interface GetToken {
  token: string
  expirationTime: string
}
export interface SysStoreType {
  menus: MenuType[]
  userInfo: Partial<UserType>
  userInfoMenus: MenuType[]
  permissionButtons: MenuType[]
  collapsed: boolean
}
export type CustomMenus = RouteRecordRaw & Partial<MenuType>

const getUserInfo = (): Promise<UserInformation | null> =>
  new Promise((resolve, reject) => {
    http<UserInformation>({
      url: 'userInfo',
      method: 'get',
    })
      .then((res) => {
        resolve(res.data)
      })
      .catch((err) => {
        reject(err)
      })
  })

async function findModName(modules: any, keyName?: string) {
  const list: any[] = []
  Object.keys(modules).forEach((key) => {
    const file = modules[key].default
    if (keyName && file.name === keyName) {
      list.push(file)
      return
    }
    list.push(file)
  })
  return list
}

async function baseLayout(item: CustomMenus[]): Promise<RouteRecordRaw[]> {
  const modules = await import.meta.globEager('../../layout/layout/**.vue')
  // 加载基础页面结构 layout
  const list: CustomMenus[] = []
  const result: any[] = await findModName(modules, 'LayoutHome')
  if (result.length > 0) {
    const data = item.filter(v => v.isHome)
    let redirect = ''
    if (data.length > 0) {
      // 优先查找是否是否存在首页，不存在直接选取第一条的路径
      redirect = String(data[0].path)
    }
    else {
      redirect = String(item[0].path)
    }

    list.push({
      path: '/',
      name: 'LayoutHome',
      component: result[0],
      id: Date.now(),
      parentId: 0,
      children: [],
      // 重定向地址
      redirect,
    })
  }
  else {
    message.error('layout文件读取失败，请检查')
  }
  return list
}

// 查询是否存在上级
function queryWhetherParent(parentId: number, menus: MenuType[]): string[] {
  let rPath: string[] = []
  const data = menus.find(v => v.id === parentId)
  if (data) {
    if (data.parentId !== 0) {
      rPath.push(data.name as string)
      const result = queryWhetherParent(data.parentId, menus)
      rPath = result.concat(rPath)
      return rPath
    }
    rPath.push(data.name as string)
    return rPath
  }
  return []
}

const formatMenuTree = async (
  menuData: MenuType[],
): Promise<RouteRecordRaw[]> => {
  // {
  //   name: 'home',
  //     path: 'home',
  //   component: () => import('@/views/home/home.vue'),
  //   meta: {
  //   title: '首页',
  //     icon: '23',
  // },
  //   redirect: '/home',
  // },
  const result: CustomMenus[] = []
  const modules = await import.meta.globEager('../../views/**/*.vue')
  const childModules = await import.meta.globEager('../../views/**/*.vue')
  menuData.forEach((menuItem) => {
    // TODO 等待支持script setup
    const fileKey = Object.keys(modules).find(
      key =>
        modules[key].default && modules[key].default.name === menuItem.name,
    )
    if (fileKey) {
      const allPath = queryWhetherParent(menuItem.parentId, menuData)
      const file = modules[fileKey].default
      const obj: any = { ...menuItem }
      obj.path = `${allPath.join('/')}/${file.name}`
      obj.name = String(menuItem.name)
      obj.component = file
      if (menuItem.name === file.name) {
        result.push(obj)
      }
      else if (menuItem.type === 1) {
        // 目录类型
        findModName(childModules).then((r) => {
          if (r && r.length > 0)
            [obj.component] = r
          result.push(obj)
        })
      }
    }
  })

  const l = await baseLayout(result)
  const [layout] = l
  layout.children = ListToTree(result)
  log.d(layout, '加载的路由')

  return [layout]
}
function listToTreeMenus(jsonData: MenuType[], parentIdItem = { id: 0, path: '', title: '' } as MenuType): MenuType[] {
  const tree: MenuType[] = []
  jsonData.forEach((item) => {
    // 每条数据的parentId 和传入的相同
    if (item.parentId === parentIdItem.id) {
      // 就去找这个元素的子集， 找到元素中parentId === item.id 递归
      // 组合路由跳转地址
      item.path = `${parentIdItem.path}${item.path}`
      if (parentIdItem.title) {
        // 组合面包屑
        item.crumb = `${parentIdItem.title},${item.title}`
      }
      else {
        item.crumb = item.title
      }
      if (parentIdItem.id)
        item.menuOpenKeys = `${parentIdItem.id},${item.id}`
      else
        item.menuOpenKeys = `${item.id}`
      const result = listToTreeMenus(jsonData, item)
      if (result && result.length > 0)
        item.children = result
      tree.push(item)
    }
  })
  return tree
}
function formatMenus(menus: MenuType[], status = false) {
  return menus.filter(item => (status ? item.type === 3 : item.type !== 3))
}
export default {
  namespace: true,
  state: {
    menus: [],
    userInfoMenus: [],
    userInfo: null,
    collapsed: false,
  },
  mutations: {
    [COLLAPSED](state: SysStoreType): void {
      state.collapsed = !state.collapsed
    },
    [SET_MENUS_MUTATION](state: SysStoreType, payload: UserInformation): void {
      const menus = formatMenus(payload.menus)
      let result: MenuType[] = []
      const list = menus.sort(ListObjCompare('orderNum'))
      list.forEach((item) => {
        if (!item.hidden) {
          result.push({
            ...item,
            key: item.id || 0,
            title: item.title,
            path: `/${String(item.name)}`,
            closable: item.isHome !== 1,
          })
        }
      })
      result = listToTreeMenus(result)
      log.w(result, 'result')
      state.menus = result
    },
    [SETUSERINFO](state: SysStoreType, payload: UserInformation): void {
      state.userInfo = payload.userInfo
    },
    [USERINFOMENUS](state: SysStoreType, payload: UserInformation): void {
      state.userInfoMenus = payload.menus
    },
    [PERMISSIONBUTTONS](state: SysStoreType, payload: UserInformation): void {
      const menus = cloneDeep(payload.menus)
      const data = formatMenus(menus, true)
      data.forEach((item) => {
        const r = menus.find(v => v.id === item.parentId)
        if (r)
          item.name = r.id
      })
      state.permissionButtons = data
    },
  },
  actions: {
    [SET_SYS]({
      commit,
    }: {
      commit: Commit
    }): Promise<{ userInfo: UserType | null; menus: RouteRecordRaw[] }> {
      return new Promise((resolve, reject) => {
        getUserInfo()
          .then(async (res) => {
            if (res) {
              // 存在选择的路由，左侧菜单对应展开
              const currentData = await storeInstant.get(CURRENT_MENU)
              if (currentData) {
                store.commit(OPEN_LEFT_MENU, currentData)
                // 面包屑默认选中
                if (currentData.crumb) {
                  log.i(currentData.crumb, 'currentData.crumb')
                  store.commit(BREAD_CRUMBS, currentData.crumb.split(','))
                }
              }
              commit(PERMISSIONBUTTONS, res)
              commit(USERINFOMENUS, res)
              commit(SETUSERINFO, res)
              commit(SET_MENUS_MUTATION, res)
              await storeInstant.set(
                LIST_OF_ALL_STORED_MENU_ITEMS,
                res.menus,
              )
              const menus = await formatMenuTree(formatMenus(res.menus))

              resolve({
                userInfo: res.userInfo,
                menus,
              })
            }
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
    [LOGIN](s: unknown, payload: LoginType): Promise<Key> {
      return new Promise<string>((resolve, reject) => {
        const body = {
          username: payload.account,
          password: payload.pass_word,
        }
        http<GetToken>({
          url: 'login',
          method: 'post',
          body,
        })
          .then(async (res) => {
            localStorage.setItem(TOKEN, res.data?.token ?? '')
            resolve(res.data?.token ?? '')
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
  },
}
