<template>
  <h1 :class="accent && 'text-blue-700'" class="font-semibold"><slot></slot></h1>
</template>

<script setup>
defineProps({
  accent: {
    type: Boolean,
    default: false
  }
})
</script>