<template>
  <NuxtLink v-if="github" :to="`https://github.com/${username}`" class="flex items-center gap-3 my-2">
    <img :src="`/${username}.webp`" class="rounded-full w-10 h-10" :alt="displayName">
    <div class="flex flex-col">
      <span class="text-sm font-semibold">{{ displayName }}</span>
      <span class="text-sm text-zinc-400">{{ task }}</span>
    </div>
  </NuxtLink>

  <div v-else class="flex items-center gap-3 my-2">
    <img :src="`/${username}.webp`" class="rounded-full w-10 h-10" :alt="displayName">
    <div class="flex flex-col">
      <span class="text-sm font-semibold">{{ displayName }}</span>
      <span class="text-sm text-zinc-400">{{ task }}</span>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  username: String,
  displayName: String,
  task: String,
  github: {
    type: Boolean,
    default: false
  }
})
</script>