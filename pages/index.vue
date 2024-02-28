<template>
  <nav class="flex justify-center py-6">
    <div class="flex justify-between items-center max-w-5xl w-full">
      <h1 class="text-2xl font-bold text-blue-700 mr-6">Unity</h1>
      <div class="flex items-center gap-1">
        <NavbarLink to="https://github.com/noahzeisberg/unity">
          <span>GitHub</span>
        </NavbarLink>

        <NavbarLink to="/stack">
          <span>Stack</span>
        </NavbarLink>

        <NavbarLink to="https://github.com/noahzeisberg/unity/issues/new">
          <span>Problem melden</span>
        </NavbarLink>
      </div>
    </div>
  </nav>

  <section class="h-screen bg-blue-700">
  </section>
</template>

<script setup>
import NavbarLink from "~/components/NavbarLink.vue";

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
