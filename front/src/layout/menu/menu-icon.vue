<template>
  <div class="icon_svg">
    <SvgIcon width="20px" height="20px" color="white" :name="name" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import SvgIcon from '@/components/SvgIcon/index.vue'

export default defineComponent({
  components: {
    SvgIcon,
  },
  props: {
    name: String,
  },
  setup() {
    return {

    }
  },
})
</script>

<style lang="less" scoped>
.icon_svg {
  margin-right: 15px;
  height: 30px;
}
</style>
