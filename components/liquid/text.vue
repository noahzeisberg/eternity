<template>
  <span :class="accent && 'text-zinc-600'"><slot></slot></span>
</template>

<script setup>
defineProps({
  accent: {
    type: Boolean,
    default: false
  }
})
</script>