<template>
  <button :class="accent && 'bg-zinc-900 ring-1 ring-zinc-800 ring-inset'" class="text-sm bg-blue-700 text-zinc-50 px-6 py-2.5 rounded-md shadow font-semibold">
    <slot></slot>
  </button>
</template>

<script setup>
defineProps({
  accent: Boolean
})
</script>