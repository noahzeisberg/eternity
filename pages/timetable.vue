<template>
  <UnityPage page-title="Stundenplan">
    <TabGroup :default-index="getCurrentDayOfWeek()">
      <TabList class="flex gap-1.5 p-1.5 m-3 rounded-lg bg-zinc-100">
        <Tab class="text-sm outline-none font-medium ui-selected:bg-white ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Mo</Tab>
        <Tab class="text-sm outline-none font-medium ui-selected:bg-white ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Di</Tab>
        <Tab class="text-sm outline-none font-medium ui-selected:bg-white ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Mi</Tab>
        <Tab class="text-sm outline-none font-medium ui-selected:bg-white ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Do</Tab>
        <Tab class="text-sm outline-none font-medium ui-selected:bg-white ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Fr</Tab>
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
  return currentDay >= 5 ? 0 : currentDay
}
</script>