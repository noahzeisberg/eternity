<template>
  <NuxtLink :to="to" class="text-zinc-700 hover:text-zinc-900"><slot></slot></NuxtLink>
</template>

<script setup>
defineProps({
  to: {
    type: String,
    required: false,
    default: "#"
  }
})
</script>