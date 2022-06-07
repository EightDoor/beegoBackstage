<template>
  <a-menu
    v-model:selectedKeys="menusInfo.selectedKeys"
    v-model:open-keys="menusInfo.openKeys"
    mode="inline"
    theme="dark"
    :inline-collapsed="collapsed"
    :collapsed="collapsed"
  >
    <template v-for="item in getMenus">
      <template v-if="!item.children">
        <a-menu-item :key="item.key" @click="jumpTo(item)">
          <div class="menu_item_container">
            <!-- 图标 -->
            <template v-if="item.icon">
              <MenuIcon :name="item.icon" />
            </template>
            <span>{{ item.title }}</span>
          </div>
        </a-menu-item>
      </template>
      <template v-else>
        <SubMenu :key="item.key" :menu-info="item" />
      </template>
    </template>
  </a-menu>
</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  onMounted,
  reactive,
  watch,
} from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import SubMenu from './menu-item.vue'
import { commSetData } from './menu-comm'
import MenuIcon from './menu-icon.vue'
import type { MenuItem, MenusInfo } from '@/types/layout/menu'

import log from '@/utils/log'

export default defineComponent({
  name: 'CommonMenu',
  components: {
    SubMenu,
    MenuIcon,
  },
  setup() {
    const router = useRouter()
    const store = useStore()
    const menusInfo = reactive<MenusInfo>({
      selectedKeys: [],
      list: [],
      openKeys: [],
    })
    const openMenus = computed(() => store.state.crumbs.menuOpenKeys)

    onMounted(() => {
      menusInfo.list = []
    })

    async function jumpTo(item: MenuItem) {
      if (item.path) {
        await commSetData(item)
        await router.push({
          path: item.path,
        })
      }
    }

    watch(openMenus, (newVal) => {
      log.i(newVal, 'newVal')
      if (newVal) {
        const result = newVal.map(item => Number(item))
        menusInfo.selectedKeys = result
        menusInfo.openKeys = result
      }
    }, {
      deep: true,
      immediate: true,
    })
    const getMenus = computed(() => store.state.sys.menus)
    const getAllMenus = computed(() => store.state.sys.userInfoMenus)

    // TODO 展开父级菜单，多级菜单有问题
    // const onOpenChange = (openKeys: number[]) => {
    //   const rootSubmenuKeys: number[] = []
    //   getAllMenus.value.forEach((item: MenuType) => {
    //     rootSubmenuKeys.push(item.id)
    //   })
    //   const latestOpenKey = openKeys.find(key => !menusInfo.openKeys.includes(key))
    //   if (!rootSubmenuKeys.includes(latestOpenKey!))
    //     menusInfo.openKeys = openKeys
    //   else
    //     menusInfo.openKeys = latestOpenKey ? [latestOpenKey] : []
    // }
    return {
      menusInfo,
      getMenus,
      collapsed: computed(() => store.state.sys.collapsed),
      jumpTo,
      // onOpenChange,
    }
  },
})
</script>

<style scoped lang="less">
@import "./menu.less";
.menu_item_container {
  display: flex;
  flex-direction: row;
  align-items: center;
}
</style>
