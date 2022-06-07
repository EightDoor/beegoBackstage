<template>
  <svg id="svg-ele" aria-hidden="true" :width="width" :height="height" viewBox="0 0 30 30">
    <use :xlink:href="symbolId" :fill="color" />
  </svg>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'

export default defineComponent({
  name: 'SvgIcon',
  props: {
    width: {
      type: String,
      default: '30px',
    },
    height: {
      type: String,
      default: '30px',
    },
    prefix: {
      type: String,
      default: 'icon',
    },
    name: {
      type: String,
      required: true,
    },
    color: {
      type: String,
      default: '#333',
    },
  },
  setup(props) {
    const symbolId = computed(() => `#${props.prefix}-${props.name}`)
    return { symbolId }
  },
})
</script>

<style lang="less" scoped>
#svg-ele {

}
</style>
