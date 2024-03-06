<template>
  <LiquidPage page-title="Eternity">
    <LiquidOutline>
      <div class="flex items-center justify-center flex-col my-6">
        <LiquidHeading class="text-2xl">{{ getDate() }}</LiquidHeading>
        <LiquidHeading>{{ getCurrentDayOfWeek() }}</LiquidHeading>
      </div>

      <LiquidCard>
        <template #header>
          <LiquidHeading>Ordnungsdienst</LiquidHeading>
        </template>

        <LiquidText>{{ selectStudents()[0] }} <LiquidText accent>(Besen)</LiquidText></LiquidText>
        <LiquidText>{{ selectStudents()[1] }} <LiquidText accent>(Kehrblech)</LiquidText></LiquidText>

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