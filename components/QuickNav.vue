<template>
  <UCard :ui="{ body: { base: 'flex justify-center items-center' } }" class="rounded-none fixed bottom-0 w-full">
    <div class="flex justify-between items-center px-3 w-full max-w-md">
      <QuickLink icon="i-heroicons-home" to="/"></QuickLink>
      <QuickLink icon="i-heroicons-calendar-days" to="/timetable"></QuickLink>
      <QuickLink icon="i-heroicons-user" to="/students"></QuickLink>
      <QuickLink icon="i-heroicons-briefcase" to="/teachers"></QuickLink>
      <QuickLink icon="i-heroicons-newspaper" to="/feed"></QuickLink>
      <QuickLink icon="i-heroicons-cog" to="/settings"></QuickLink>
    </div>
  </UCard>

  <UContainer class="h-16"></UContainer>
</template>
<script setup lang="ts">
</script>