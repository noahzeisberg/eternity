<template>
  <button @click="makeApiCall">Test</button>
  <input v-model="number" type="number">
  <span>{{ text }}</span>
</template>

<script setup>
const text = ref("No information")
const number = ref(0)

async function makeApiCall() {
  const data = await $fetch("/api/students", {
    method: "GET"
  })

  text.value = data["students"][number.value]
}
</script>