<template>
  <input :placeholder="placeholder" class="text-sm px-4 py-2.5 ring-1 ring-zinc-800 ring-inset rounded-md shadow bg-zinc-900 placeholder-zinc-600 outline-none focus:outline-2 focus:outline-blue-700 duration-300 ease-in-out" :type="type" v-model="model">
</template>

<script setup>
defineProps(["placeholder", "type"])
const model = defineModel()
</script>