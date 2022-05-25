<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  isRouter: true,
  name: 'LayoutChildren',
})
</script>

<template>
  <router-view />
</template>
