<template>
  <BaseContainer>
    <fs-crud ref="crudRef" v-bind="crudBinding" />
  </BaseContainer>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useCrud, useExpose } from '@fast-crud/fast-crud'
import createCrudOptions from './crud'
import BaseContainer from '@/components/BaseContainer/index.vue'

export default defineComponent({
  name: 'DemoAvue', // 实际开发中可以修改一下name
  components: {
    BaseContainer,
  },
  setup() {
    // crud组件的ref
    const crudRef = ref<any>()
    // crud 配置的ref
    const crudBinding = ref<any>()
    // 暴露的方法
    const { crudExpose } = useExpose({ crudRef, crudBinding })

    // 你的crud配置
    const { crudOptions } = createCrudOptions({ crudExpose })
    // 初始化crud配置
    const { resetCrudOptions } = useCrud({ crudExpose, crudOptions })
    // 你可以调用此方法，重新初始化crud配置
    // resetCrudOptions(options)
    // 页面打开后获取列表数据
    onMounted(() => {
      crudExpose.doRefresh()
    })
    return {
      crudBinding,
      crudRef,
    }
  },
})

function crudExpose(crudExpose: any): { crudOptions: any } {
  throw new Error('Function not implemented.')
}
</script>
