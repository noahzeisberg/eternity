<template>
  <UCard class="rounded-none">
    <div class="flex justify-between items-center -my-2">
      <UButton to="/settings" color="gray" variant="ghost" icon="i-heroicons-cog"></UButton>
      <h1 class="font-semibold">{{ title }}</h1>
      <UButton @click="isOpen = true" color="gray" variant="ghost" icon="i-heroicons-bars-3"></UButton>
    </div>
  </UCard>

  <main class="flex flex-col">
    <slot></slot>

    <!-- Bottom navigation bar left in for future reimplementation. -->

    <UCard v-if="false" class="rounded-none fixed bottom-0 w-full">
      <div class="flex justify-between items-center px-5">
        <ULink class="size-6" to="/"><UIcon name="i-heroicons-home-solid" class="size-6"></UIcon></ULink>
        <ULink class="size-6" to="/"><UIcon name="i-heroicons-star-solid" class="size-6"></UIcon></ULink>
      </div>
    </UCard>
  </main>

  <Sidebar v-model="isOpen"></Sidebar>
  <UNotifications></UNotifications>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    default: "Eternity"
  }
})
const isOpen = ref(false)
</script>