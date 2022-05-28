<template>
  <a-modal
    v-model:visible="commdrawerData.visible"
    :title="title"
    :width="width"
    destroy-on-close
    :footer="props.footer ? undefined : null"
    @cancel="onCancel"
  >
    <div class="drawerContainer">
      <slot />
    </div>
    <template #footer>
      <div v-if="props.footer">
        <a-button key="back" :loading="commdrawerData.loading" styl="margin-right: 15px" @click="onCancel">
          取消
        </a-button>
        <a-button key="submit" :loading="commdrawerData.loading" type="primary" @click="onOk">
          确定
        </a-button>
      </div>
    </template>
  </a-modal>
</template>

<script lang="ts">
import { defineComponent, h, reactive, watch } from 'vue'
import { Button } from 'ant-design-vue'
import log from '@/utils/log'

export interface DrawerProps {
  visible: boolean
  title: string
  placement?: string
  loading?: boolean
}

const CommonDrawer = defineComponent({
  name: 'ComponentDrawer',
  props: {
    title: {
      type: String,
      required: true,
    },
    placement: {
      type: String,
      default: 'right',
    },
    okText: {
      type: String,
      default: '',
    },
    cancelText: {
      type: String,
      default: '',
    },
    width: {
      type: String,
      default: '45%',
    },
    visible: {
      type: Boolean,
      default: false,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    footer: {
      type: Boolean,
      default: true,
    },
  },
  emits: ['on-close', 'on-ok'],
  setup(props, { emit }) {
    const commdrawerData = reactive({
      visible: false,
      loading: false,
    })
    function onCancel() {
      emit('on-close')
      commdrawerData.visible = false
    }
    function onOk() {
      emit('on-ok')

      // 超时自定关闭loading
      setTimeout(() => {
        commdrawerData.loading = false
      }, 5000)
    }
    function onClose() {
      emit('on-close')
      commdrawerData.visible = false
    }
    watch(
      () => props.visible,
      (data) => {
        log.d(data, 'Modal - Data')
        commdrawerData.visible = data
        commdrawerData.loading = false
      },
      {
        immediate: true,
        deep: true,
      },
    )
    return {
      commdrawerData,
      onCancel,
      onOk,
      onClose,
      props,
    }
  },
})

export default CommonDrawer
</script>

<style lang="less" scoped>
@import './Drawer.less';
</style>
