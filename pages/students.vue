<template>
  <Page title="Schüler">
    <Loader v-if="pending"></Loader>

    <UTable class="w-full" v-else :rows="rows"></UTable>
  </Page>
</template>

<script setup lang="ts">
const { pending, data: students } = useLazyAsyncData("students", () => $fetch("/api/students"))
const rows = []

watch(students, (list) => {
  if(list === null) return
  list.forEach((student) => {
    rows.push({
      name: student
    })
  })
})
</script>