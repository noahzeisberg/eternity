<template>
  <UnityPage page-title="Stundenplan">
    <TabGroup :default-index="getCurrentDayOfWeek()">
      <TabList class="flex justify-stretch gap-3 w-full overflow-y-scroll p-3">
        <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2 ui-selected:bg-blue-50 ui-selected:text-blue-700">Montag</Tab>
        <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2 ui-selected:bg-blue-50 ui-selected:text-blue-700">Dienstag</Tab>
        <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2 ui-selected:bg-blue-50 ui-selected:text-blue-700">Mittwoch</Tab>
        <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2 ui-selected:bg-blue-50 ui-selected:text-blue-700">Donnerstag</Tab>
        <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2 ui-selected:bg-blue-50 ui-selected:text-blue-700">Freitag</Tab>
      </TabList>
      <TabPanels>
        <TabPanel v-for="day in timetable">
          <Outline>
            <Card v-for="(lesson, index) in day.lessons">
              <template #header>
                <h1 class="font-semibold">{{ index+1 }}. {{ lesson.name }}</h1>
              </template>

              <p>Lehrer: {{ getTeacherByShort(lesson.teacher).name }} <NuxtLink :to="`/teacher/${lesson.teacher}`" class="text-zinc-400 underline">({{ lesson.teacher }})</NuxtLink></p>
              <p>Raum: {{ lesson.room }}</p>
            </Card>
          </Outline>
        </TabPanel>
      </TabPanels>
    </TabGroup>
  </UnityPage>
</template>
<script setup>
import {Tab, TabGroup, TabList, TabPanel, TabPanels} from "@headlessui/vue";
const { timetable } = await getTimetable()
const { teachers } = await getTeachers()

function getTeacherByShort(short) {
  for (let i = 0; i < teachers.length; i++) {
    if(teachers[i].short === short) {
      return teachers[i]
    }
  }
  return {}
}

async function getTimetable() {
  const { timetable } = await $fetch("/api/timetable")
  return { timetable }
}

async function getTeachers() {
  const { teachers } = await $fetch("/api/teachers")
  return { teachers }
}

function getCurrentDayOfWeek() {
  const today = new Date().getDay();
  const currentDay = (today + 6) % 7
  const selected = currentDay >= 5 ? 0 : currentDay
  return selected
}
</script>