<template>
  <ULink :to="`mailto://${email}`" class="flex justify-start items-center gap-4 py-2">
    <UAvatar size="md" :src="`/${img}.webp`" :alt="name"></UAvatar>
    <div class="flex flex-col">
      <h1 class="font-semibold">{{ name }}</h1>
      <h1 class="text-sm text-zinc-400">{{ task }}</h1>
    </div>
  </ULink>
</template>

<script setup lang="ts">
defineProps({
  name: String,
  img: String,
  email: String,
  task: String,
})
</script>