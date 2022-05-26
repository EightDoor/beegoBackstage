<script lang="ts">
import {
  computed,
  defineComponent, ref, toRaw, watch,
} from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import localStore from '@/utils/store'
import { CURRENT_MENU } from '@/utils/constant'
import type { SysTab, SysTabDel } from '@/types/sys/tab'
import { DELETETABS, DELETETABSACTION } from '@/store/mutation-types'
import log from '@/utils/log'

export default defineComponent({
  name: 'TabsHome',
  setup() {
    const store = useStore()
    const router = useRouter()
    const activeKey = ref(0)

    const panes = ref<{
      title?: string
      closable?: boolean
    }[]>([])

    const pansList = computed(() => store.state.crumbs.panes)
    watch(
      pansList,
      (newVal) => {
        log.i(newVal, '更新的值')
        localStore.get<SysTab>(CURRENT_MENU).then((res) => {
          if (newVal && newVal.length > 0) {
            log.i(res, '获取到的菜单')
            log.i(newVal, 'newVal')
            if (res) {
              const result = newVal.findIndex(item => item.id === res.id)
              // 设置当前被选择项
              if (result !== -1)
                activeKey.value = result
              else
                activeKey.value = newVal.length - 1
            }
            panes.value = newVal
          }
        })
      },
      {
        deep: true,
        immediate: true,
      },
    )

    const FormatData = computed(() => store.state.crumbs.panes)
    const OnChange = (val: number) => {
      const result = FormatData.value[val]
      activeKey.value = val
      router.push(result.path)
      store.commit(DELETETABS, { selectData: toRaw(result) })
    }

    const OnEdit = (val: number) => {
      log.i(val, 'val')
      const value = activeKey.value
      const result: SysTabDel = {
        delData: FormatData.value[val],
        selectData: '',
      }
      const len = FormatData.value.length
      if (len > 0) {
        const index = value >= val ? value - 1 : value
        result.selectData = FormatData.value[index]
      }
      store.dispatch(DELETETABSACTION, result)
    }
    return {
      activeKey,
      OnChange,
      panes,
      OnEdit,
    }
  },
})
</script>

<template>
  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      hide-add
      type="editable-card"
      @change="OnChange"
      @edit="OnEdit"
    >
      <a-tab-pane
        v-for="(pane, index) in panes"
        :key="index"
        :tab="pane.title"
        :closable="pane.closable"
      >
        <slot />
      </a-tab-pane>
      <template #rightExtra>
        <a-button style="padding-right: 15px" type="primary">
          关闭全部
        </a-button>
      </template>
    </a-tabs>
  </div>
</template>

<style lang="less" scoped>
@import "index.less";
</style>
