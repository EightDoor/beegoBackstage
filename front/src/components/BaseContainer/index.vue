<template>
  <div class="card_container">
    <slot />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'BaseContainer',
})
</script>

<style lang="less">
.card_container {
  //height: 80vh;
  //overflow-y: auto;
}
</style>
