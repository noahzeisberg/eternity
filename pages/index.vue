<template>
  <Page page-title="Eternity">
    <Outline>
      <div class="flex items-center justify-center flex-col my-3">
        <h1 class="text-2xl font-semibold">{{ getDate() }}</h1>
        <h1>{{ getCurrentDayOfWeek() }}</h1>
      </div>

      <Card>
        <template #header>
          <h1 class="font-semibold">Ordnungsdienst</h1>
        </template>

        <span>{{ selectStudents()[0] }} <span class="text-zinc-400">(Besen)</span></span>
        <span>{{ selectStudents()[1] }} <span class="text-zinc-400">(Kehrblech)</span></span>

        <template #footer>
          <span class="text-sm text-zinc-500">Zyklus: {{ cycle }} - {{ week }}</span>
        </template>
      </Card>
    </Outline>
  </Page>
</template>

<script setup>
let cycle
const week = getCurrentWeekOfYear()
const { students } = await $fetch("/api/students")

function selectStudents() {
  let targetLength
  let groups = []
  if(students.length % 2 === 0) {
    targetLength = students.length
  } else {
    students.push(...students)
    targetLength = students.length*2
  }

  let even = []
  let odd = []

  for (let i = 0; i < targetLength; i++) {
    if(i % 2 === 0) {
      even.push(students[i])
    } else {
      odd.push(students[i])
    }
  }

  for (let i = 0; i < even.length; i++) {
    groups.push({
      0: even[i],
      1: odd[i]
    })
  }

  if (week >= groups.length) {
    cycle = week % groups.length
  } else {
    cycle = week
  }

  return groups[cycle]
}

function getCurrentWeekOfYear() {
  const now = new Date()
  const startOfYear = new Date(now.getFullYear(), 0, 0)
  const diff = now - startOfYear
  const oneWeekInMillis = 1000 * 60 * 60 * 24 * 7
  return Math.floor(diff / oneWeekInMillis)
}

function getDate() {
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
}

function getCurrentDayOfWeek() {
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