<template>
  <Page title="Schüler">
    <Loader v-if="pending"></Loader>

    <Wrapper v-else>
      <UCard v-for="student in students">
        {{ student }}
      </UCard>
    </Wrapper>
  </Page>
</template>

<script setup lang="ts">
const { pending, data: students } = useLazyAsyncData("students", () => $fetch("/api/students"))
</script>