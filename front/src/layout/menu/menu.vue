<script lang="ts">
import {
  computed,
  defineComponent,
  onMounted,
  reactive,
  toRaw,
} from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import SubMenu from './menu-item.vue'
import { commSetData } from './menu-comm'
import type { MenuItem, MenusInfo } from '@/types/layout/menu'
import type { MenuType } from '@/types/sys'

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

    async function jumpTo(item: MenuItem) {
      if (item.path) {
        await commSetData(item)
        router.push({
          path: item.path,
        })
      }
    }
    return {
      menusInfo,
      getMenus: computed(() => store.state.sys.menus),
      collapsed: computed(() => store.state.sys.collapsed),

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
