<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { commSetData } from './menu-comm'
import type { MenuItem } from '@/types/layout/menu'

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
        await commSetData(item)
        await router.push({
          path: item.path,
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
