<template>
  <ULink class="flex justify-center items-center size-6" :to="to">
    <UIcon dynamic :name="route === to ? icon+'-solid' : icon" class="size-6"></UIcon>
  </ULink>
</template>

<script setup lang="ts">
const route = useRoute().path
defineProps({
  to: String,
  icon: String
})
</script>