<template>
  <NuxtLink class="flex gap-3 items-center hover:bg-zinc-100 hover:dark:bg-zinc-900 px-5 py-1.5 rounded-md" active-class="bg-zinc-100 dark:bg-zinc-900" :to="to">
    <slot></slot>
    <h1>{{ name }}</h1>
  </NuxtLink>
</template>

<script setup lang="ts">
defineProps({
  to: String,
  name: String
})
</script>