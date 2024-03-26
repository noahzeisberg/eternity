<template>
  <Page title="Arbeitsplan">
    <Loader v-if="pending"></Loader>

    <UTabs class="w-full" v-else :items="parsedExams">
      <template #item="{ item }">
        <div class="w-full flex items-center justify-center">
          <Wrapper>
            <UCard v-for="exam in item.content">
              <template #header>
                <div class="flex justify-between items-center">
                  <h1 class="font-semibold -my-2">{{ exam.name }}</h1>
                  <h1 class="text-sm text-zinc-400 -my-2">{{ exam.date }}. {{ item.label }}</h1>
                </div>
              </template>
              <div>Lehrer: {{ exam.teacher.toString().replace(",", ", ") }}</div>
            </UCard>
          </Wrapper>
        </div>
      </template>
    </UTabs>
  </Page>
</template>

<script setup lang="ts">
const { pending, data: exams } = await useLazyAsyncData("exams", () => $fetch("/api/exams"))
const parsedExams: {label: string, content: any}[] = []
watch(exams, (table) => {
  if (table === null) return
  table.forEach(item => {
    if(item.exams.length === 0) return
    parsedExams.push({
      label: item.name,
      content: item.exams
    })
  })
})
</script>