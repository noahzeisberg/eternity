<template>
  <NuxtLink class="text-zinc-950 flex items-center gap-2 p-2 rounded-lg hover:bg-blue-50 hover:text-blue-700 duration-300 ease-in-out" :to="to"><slot></slot></NuxtLink>
</template>

<script setup>
defineProps({
  to: {
    required: false,
    default: "/",
    type: String
  }
})
</script>