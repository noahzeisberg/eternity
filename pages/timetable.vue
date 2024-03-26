<template>
  <Page title="Stundenplan">
    <Loader v-if="pending"></Loader>

    <UTabs class="w-full" :ui="{ list: { tab: { font: 'font-medium' } } }" v-else :items="parsedTimetable">
      <template #item="{ item }">
        <div class="w-full flex items-center justify-center">
          <Wrapper>
            <UCard v-for="(lesson, index) in item.content">
              <template #header>
                <h1 class="font-semibold -my-2">{{ index+1 }}. {{ lesson.name }}</h1>
              </template>
              <div>Lehrer: {{ lesson.teacher }}</div>
              <div>Raum: {{ lesson.room }}</div>
            </UCard>
          </Wrapper>
        </div>
      </template>
    </UTabs>
  </Page>
</template>

<script setup lang="ts">
const { pending, data: timetable } = await useLazyAsyncData("timetable", () => $fetch('/api/timetable'))
let parsedTimetable: {label: string, content: any}[] = []

watch(timetable, (table) => {
  if (table === null) return
  table.forEach(item => {
    parsedTimetable.push({
      label: item.day.substring(0, 2),
      content: item.lessons
    })
  })
})
</script>