<script lang="ts">
import {
  computed,
  defineComponent,
  onMounted,
  reactive,
  toRaw, watch,
} from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import SubMenu from './menu-item.vue'
import type { MenuItem, MenusInfo } from '@/types/layout/menu'
import type { MenuType } from '@/types/sys'
import log from '@/utils/log'

interface InitTopTabs extends MenuItem {
  crumbs: string
}
export default defineComponent({
  name: 'CommonMenu',
  components: {
    SubMenu,
  },
  setup() {
    const router = useRouter()
    const store = useStore()
    const menusInfo = reactive<MenusInfo>({
      selectedKeys: [],
      list: [],
      openKeys: [],
    })
    const getUserInfoMenus = computed(() => store.state.sys.userInfoMenus)

    const FormatSelectKey = (res) => {
      log.i(res, 'FormatSelectKey - res')
      menusInfo.selectedKeys = [res.key || res.id || '']
      const { parentId } = res
      const data: MenuType[] = toRaw(getUserInfoMenus.value)
      const r = data.find(item => item.id === parentId)
      if (r)
        menusInfo.openKeys = [r.id]
    }

    onMounted(() => {
      menusInfo.list = []
    })

    // methods
    function jumpTo(item: MenuItem) {
      if (item.path) {
        router.push({
          path: item.path || '',
        })
      }
    }
    return {
      // data
      menusInfo,
      getMenus: computed(() => store.state.sys.menus),
      collapsed: computed(() => store.state.sys.collapsed),
      // methods
      jumpTo,
    }
  },
})
</script>

<template>
  <a-menu
    v-model:openKeys="menusInfo.openKeys"
    v-model:selectedKeys="menusInfo.selectedKeys"
    mode="inline"
    theme="dark"
    :inline-collapsed="collapsed"
    :collapsed="collapsed"
  >
    <template v-for="item in getMenus">
      <template v-if="!item.children">
        <a-menu-item :key="item.key" @click="jumpTo(item)">
          <!-- 图标 -->
          <span>{{ item.title }}</span>
        </a-menu-item>
      </template>
      <template v-else>
        <SubMenu :key="item.key" :menu-info="item" />
      </template>
    </template>
  </a-menu>
</template>

<style scoped lang="less">
@import "./menu.less";
</style>
