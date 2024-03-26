<template>
  <UCard class="rounded-none">
    <div class="flex justify-between items-center -my-2">
      <ULink to="/settings"><UAvatar size="sm" :alt="useCookie('user.name').value"></UAvatar></ULink>
      <h1 class="font-semibold">{{ title }}</h1>
      <UButton @click="isOpen = true" color="gray" variant="ghost" icon="i-heroicons-bars-3"></UButton>
    </div>
  </UCard>

  <main class="flex flex-col">
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