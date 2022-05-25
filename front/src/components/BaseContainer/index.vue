<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'BaseContainer',
})
</script>

<template>
  <div class="card_container">
    <slot />
  </div>
</template>

<style lang="less">
.card_container {
  height: 80vh;
  border: 1px solud red;
}
</style>
