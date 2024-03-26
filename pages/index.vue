<template>
  <Page title="Eternity">
    <Wrapper>
      <UContainer class="flex flex-col items-center justify-center gap-1 my-6">
        <h1 class="text-3xl font-semibold">{{ getCurrentDayOfWeek() }}</h1>
        <h1>{{ getDate() }}</h1>
      </UContainer>

      <UCard>
        <template #header>
          <h1 class="font-semibold -my-2">Ordnungsdienst</h1>
        </template>

        <Loader v-if="pending"></Loader>

        <div v-else class="flex flex-col">
          <span>{{ student1 }}</span>
          <span>{{ student2 }}</span>
        </div>
      </UCard>
    </Wrapper>
  </Page>
</template>

<script setup lang="ts">
const student1 = ref("")
const student2 = ref("")

let cycle
const getCurrentWeek = () => {
  const now = new Date()
  const startOfYear = new Date(now.getFullYear(), 0, 0)
  const diff = now.getTime() - startOfYear.getTime()
  const oneWeekInMillis = 1000 * 60 * 60 * 24 * 7
  return Math.floor(diff / oneWeekInMillis)
}

const { pending, data: students } = await useLazyAsyncData("students", () => $fetch("/api/students"))
watch(students, (list) => {
  if(list === null) return
  let targetLength
  let groups = []
  if(list.length % 2 === 0) {
    targetLength = list.length
  } else {
    list.push(...list)
    targetLength = list.length
  }

  let even = []
  let odd = []

  for (let i = 0; i < targetLength; i++) {
    if(i % 2 === 0) {
      even.push(list[i])
    } else {
      odd.push(list[i])
    }
  }

  for (let i = 0; i < even.length; i++) {
    groups.push({
      0: even[i],
      1: odd[i]
    })
  }

  if (getCurrentWeek() >= groups.length) {
    cycle = getCurrentWeek() % groups.length
  } else {
    cycle = getCurrentWeek()
  }

  student1.value = groups[cycle][0]
  student2.value = groups[cycle][1]
})

const getDate = () => {
  const date = new Date()
  return date.getDate() + ". " + [
    "Januar",
    "Februar",
    "März",
    "April",
    "Mai",
    "Juni",
    "Juli",
    "August",
    "September",
    "Oktober",
    "November",
    "Dezember"
  ][date.getMonth()] + " " + date.getFullYear()
};

const getCurrentDayOfWeek = () => {
  const today = new Date().getDay();
  const currentDay = (today + 6) % 7
  return [
    "Montag",
    "Dienstag",
    "Mittwoch",
    "Donnerstag",
    "Freitag",
    "Samstag",
    "Sonntag",
  ][currentDay]
}
</script>