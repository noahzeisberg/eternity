<template>
  <UCard class="rounded-none">
    <div class="flex justify-between items-center -my-2">
      <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="size-8">
        <path fill-rule="evenodd" clip-rule="evenodd" d="M11 8L3 16H11L11 8ZM13 16L21 8L13 8V16Z" fill="currentColor"/>
      </svg>
      <h1 class="font-semibold">{{ title }}</h1>
      <UButton @click="isOpen = true" color="gray" variant="ghost" icon="i-heroicons-bars-3"></UButton>
    </div>
  </UCard>

  <main class="flex flex-col items-center">
    <slot></slot>

    <QuickNav></QuickNav>
  </main>

  <Sidebar v-model="isOpen"></Sidebar>

  <UNotifications></UNotifications>
  <UModals></UModals>
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