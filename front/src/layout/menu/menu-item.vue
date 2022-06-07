<template>
  <a-sub-menu :key="menuInfo.key" v-bind="$attrs">
    <template #title>
      <div>
        <!-- 图标 -->
        <div class="menu_item_container">
          <template v-if="menuInfo.icon">
            <MenuIcon :name="menuInfo.icon" />
          </template>
          <span>{{ menuInfo.title }}</span>
        </div>
      </div>
    </template>
    <template v-for="item in menuInfo.children">
      <template v-if="!item.children">
        <a-menu-item :key="item.key" @click="jumpTo(item)">
          <!-- 图标 -->
          <div class="menu_item_container">
            <MenuIcon v-if="item.icon" :name="item.icon" />
            <span>{{ item.title }}</span>
          </div>
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.key" :menu-info="item" />
      </template>
    </template>
  </a-sub-menu>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { commSetData } from './menu-comm'
import MenuIcon from './menu-icon.vue'
import type { MenuItem } from '@/types/layout/menu'

export default defineComponent({
  name: 'SubMenu',
  components: {
    MenuIcon,
  },
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

<style lang="less" scoped>
.menu_item_container {
  display: flex;
  flex-direction: row;
  align-items: center;
}
</style>
