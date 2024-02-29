<template>
  <div class="flex flex-col gap-6 max-w-xl w-full h-auto pb-6">
    <Card>
      <h1 class="text-2xl font-bold my-2">Ordnungsdienst</h1>
      <span>{{ student1 }}</span>
      <span>{{ student2 }}</span>

      <template #footer>
        <span class="text-sm text-zinc-500">Zyklus: {{ currentCycle }} - {{ week }}</span>
      </template>
    </Card>
  </div>
</template>

<script setup>
const student1 = ref("")
const student2 = ref("")

let currentCycle
const week = getCurrentWeekOfYear()
const { students, length } = await getStudents()
const studentCount = length

currentCycle = week > studentCount ? week - studentCount : week
student1.value = students[currentCycle]

const midIndex = Math.floor(currentCycle + studentCount / 2)
student2.value = students[midIndex % studentCount]

function getCurrentWeekOfYear() {
  const now = new Date()
  const startOfYear = new Date(now.getFullYear(), 0, 0)
  const diff = now - startOfYear
  const oneWeekInMillis = 1000 * 60 * 60 * 24 * 7
  return Math.floor(diff / oneWeekInMillis)
}

async function getStudents() {
  const { students } = await $fetch("/api/students")
  return { students, length: students.length }
}
</script>