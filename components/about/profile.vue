<template>
  <div class="flex items-center gap-3 my-2">
    <img :src="`/${picture}.webp`" class="rounded-full w-10 h-10" :alt="displayName">
    <div class="flex flex-col">
      <span>{{ displayName }}</span>
      <span class="text-sm text-zinc-400">{{ task }}</span>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  picture: String,
  displayName: String,
  task: String
})
</script>