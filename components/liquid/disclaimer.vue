<template>
  <span :class="accent && 'text-zinc-400'" class="text-sm"><slot></slot></span>
</template>

<script setup>
defineProps({
  accent: {
    type: Boolean,
    default: false
  }
})
</script>