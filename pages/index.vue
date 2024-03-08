<template>
  <LiquidPage page-title="Eternity">
    <LiquidOutline>
      <div class="flex items-center justify-center flex-col my-6">
        <LiquidTitle class="text-2xl">{{ getDate() }}</LiquidTitle>
        <LiquidTitle>{{ getCurrentDayOfWeek() }}</LiquidTitle>
      </div>

      <LiquidCard>
        <template #header>
          <LiquidTitle>Ordnungsdienst</LiquidTitle>
        </template>

        <LiquidText>{{ student1 }} <LiquidText accent>(Besen)</LiquidText></LiquidText>
        <LiquidText>{{ student2 }} <LiquidText accent>(Kehrblech)</LiquidText></LiquidText>

        <template #footer>
          <LiquidDisclaimer accent>Zyklus: {{ cycle }} - {{ week }}</LiquidDisclaimer>
        </template>
      </LiquidCard>
    </LiquidOutline>
  </LiquidPage>
</template>

<script setup>
let cycle
const week = getCurrentWeekOfYear()
const { students } = await $fetch("/api/students")

const student1 = ref("")
const student2 = ref("")

selectStudents()

function selectStudents() {
  let targetLength
  let groups = []
  if(students.length % 2 === 0) {
    targetLength = students.length
  } else {
    students.push(...students)
    targetLength = students.length
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

  student1.value = groups[cycle][0]
  student2.value = groups[cycle][1]
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