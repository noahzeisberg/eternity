<template>
  <NuxtLink class="flex flex-col items-center justify-center gap-0.5" :to="to">
    <slot></slot>
  </NuxtLink>
</template>

<script setup>
const props = defineProps({
  to: String
})
</script>