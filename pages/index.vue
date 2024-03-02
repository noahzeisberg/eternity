<template>
  <UnityPage page-title="Unity">
    <Outline>
      <Card>
        <template #header>
          <h1 class="font-semibold">Übersicht</h1>
        </template>

        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Adipisci autem consequuntur delectus, facilis ipsa, laborum numquam odit quis reprehenderit vel voluptate voluptatum. Accusamus debitis et illum libero, neque odit sint?
      </Card>

      <Card>
        <template #header>
          <h1 class="font-semibold">Ordnungsdienst</h1>
        </template>

        <span>{{ student1 }} <span class="text-zinc-400">(Besen)</span></span>
        <span>{{ student2 }} <span class="text-zinc-400">(Kehrblech)</span></span>

        <template #footer>
          <span class="text-sm text-zinc-500">Zyklus: {{ currentCycle }} - {{ week }}</span>
        </template>
      </Card>
    </Outline>
  </UnityPage>
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