<template>
  <div class="container">
    <!--    <a-tabs -->
    <!--      v-model:activeKey="activeKey" -->
    <!--      hide-add -->
    <!--      type="editable-card" -->
    <!--      @change="OnChange" -->
    <!--      @edit="OnEdit" -->
    <!--    > -->
    <!--      <a-tab-pane -->
    <!--        v-for="(pane, index) in panes" -->
    <!--        :key="index" -->
    <!--        :tab="pane.title" -->
    <!--        :closable="pane.closable" -->
    <!--      > -->
    <!--        <slot /> -->
    <!--      </a-tab-pane> -->
    <!--      <template #rightExtra> -->
    <!--        <a-button style="padding-right: 15px" type="primary" @click="closeAllTab"> -->
    <!--          关闭全部 -->
    <!--        </a-button> -->
    <!--      </template> -->
    <!--    </a-tabs> -->
  </div>
</template>

<script lang="ts">
import {
  createVNode,
  defineComponent, ref,
} from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { Modal } from 'ant-design-vue'
import { ExclamationCircleOutlined } from '@ant-design/icons-vue'

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

    const OnChange = (val: number) => {
      activeKey.value = val
    }

    const OnEdit = (val: number) => {
      log.i(val, 'val')
      const value = activeKey.value
    }

    async function closeAllTab() {
      const model = Modal.confirm({
        title: '确定关闭全部吗?',
        icon: createVNode(ExclamationCircleOutlined),
        async onOk() {
          console.log('OK')
          model.destroy()
        },
        onCancel() {
          console.log('Cancel')
        },
        class: 'test',
      })
    }
    return {
      activeKey,
      OnChange,
      panes,
      OnEdit,
      closeAllTab,
    }
  },
})
</script>

<style lang="less" scoped>
@import "index.less";
</style>
