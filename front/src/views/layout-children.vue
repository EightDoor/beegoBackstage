<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  isRouter: true,
  name: 'LayoutChildren',
})
</script>
