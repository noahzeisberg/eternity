<template>
  <NuxtLink :to="link" class="flex items-center gap-3 p-2 rounded-lg hover:bg-zinc-100 hover:dark:bg-zinc-900">
    <img alt="" :src="'/' + name.toLowerCase() + '.webp'" class="size-10 rounded-full bg-zinc-200">
    <div class="flex flex-col">
      <h1 class="text-sm font-bold">{{ name }}</h1>
      <p class="text-sm text-zinc-500 dark:text-zinc-400">{{ role }}</p>
    </div>
  </NuxtLink>
</template>

<script setup>
defineProps({
  name: String,
  role: String,
  link: String,
})
</script>