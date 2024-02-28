<template>
  <div class="flex flex-col">
    <h1>Diese Woche:</h1>
    <span>{{ student1 }}</span>
    <span>{{ student2 }}</span>
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
  const { students } = await $fetch("/api/students", {
    method: "GET"
  })
  return { students, length: students.length }
}
</script>
