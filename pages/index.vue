<template>
  <UnityPage page-title="Unity">
    <Outline>
      <div class="flex items-center justify-center flex-col my-3">
        <h1 class="text-2xl font-semibold">{{ date }}</h1>
        <h1>{{ getCurrentDayOfWeek() }}</h1>
      </div>

      <Card>
        <template #header>
          <h1 class="font-semibold">Ordnungsdienst</h1>
        </template>

        <span>{{ student1 }} <span class="text-zinc-400">(Besen)</span></span>
        <span>{{ student2 }} <span class="text-zinc-400">(Kehrblech)</span></span>

        <template #footer>
          <span class="text-sm text-zinc-500">Zyklus: {{ cycle }} - {{ week }}</span>
        </template>
      </Card>
    </Outline>
  </UnityPage>
</template>

<script setup>
const student1 = ref("")
const student2 = ref("")

const date = ref(getDate())
let targetLength

const week = getCurrentWeekOfYear()
let cycle
const { students, studentCount } = await getStudents()
const selectedStudents = selectStudents()

if (week >= selectedStudents.length) {
  cycle = week % selectedStudents.length
} else {
  cycle = week
}

student1.value = selectedStudents[cycle][0]
student2.value = selectedStudents[cycle][1]

function selectStudents() {
  let groups = []
  if(studentCount % 2 === 0) {
    targetLength = studentCount
  } else {
    students.push(...students)
    targetLength = studentCount*2
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

  return groups
}

function getCurrentWeekOfYear() {
  const now = new Date()
  const startOfYear = new Date(now.getFullYear(), 0, 0)
  const diff = now - startOfYear
  const oneWeekInMillis = 1000 * 60 * 60 * 24 * 7
  return Math.floor(diff / oneWeekInMillis)
}

async function getStudents() {
  const { students } = await $fetch("/api/students")
  return { students, studentCount: students.length }
}

function getDate() {
  const date = new Date()
  return date.getDay() + ". " + [
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
  ][date.getMonth()-1] + " " + date.getFullYear()
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