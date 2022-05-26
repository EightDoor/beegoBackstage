<script lang="ts">
import { defineComponent, toRaw } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { MenuFormatBrumb } from './menu-common'
import type { MenuItem } from '@/types/layout/menu'
import { CURRENT_MENU, STORELETMENUPATH } from '@/utils/constant'
import localStoreInstant from '@/utils/store'
import { SETCRUMBSLIST } from '@/store/mutation-types'
import log from '@/utils/log'
import { formatArr } from '@/utils'

export default defineComponent({
  name: 'SubMenu',
  props: {
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  setup() {
    const router = useRouter()
    const store = useStore()
    // methods
    async function jumpTo(item: MenuItem) {
      if (item.path) {
        store.commit(SETCRUMBSLIST, toRaw(item.crumbs))
        await localStoreInstant.set(CURRENT_MENU, toRaw(item))
        localStoreInstant.get(STORELETMENUPATH).then((res) => {
          log.i(res, '点击二级菜单-获取的存储值')
          let data: MenuItem[] = []
          if (res)
            data = [...res, toRaw(item)]
          else
            data = [toRaw(item)]

          localStoreInstant.set(STORELETMENUPATH, formatArr(data)).then(() => {
            router.push({
              path: item.path || '',
            })
            MenuFormatBrumb(item)
          })
        })
      }
    }
    return {
      // methods
      jumpTo,
    }
  },
})
</script>

<template>
  <a-sub-menu :key="menuInfo.key" v-bind="$attrs">
    <template #title>
      <span>
        <!-- 图标 -->
        {{ menuInfo.title }}</span>
    </template>
    <template v-for="item in menuInfo.children">
      <template v-if="!item.children">
        <a-menu-item :key="item.key" @click="jumpTo(item)">
          <!-- 图标 -->
          <span>{{ item.title }}</span>
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.key" :menu-info="item" />
      </template>
    </template>
  </a-sub-menu>
</template>
