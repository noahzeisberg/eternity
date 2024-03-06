<template>
  <LiquidPage page-title="Stundenplan">
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
          <LiquidOutline>
            <LiquidCard v-for="(lesson, index) in day.lessons">
              <template #header>
                <LiquidTitle>{{ index+1 }}. {{ lesson.name }}</LiquidTitle>
              </template>

              <LiquidText>Lehrer: {{ getTeacherByShort(lesson.teacher).name }} <NuxtLink :to="`/teacher/${lesson.teacher}`" class="text-zinc-400 underline">({{ lesson.teacher }})</NuxtLink></LiquidText>
              <LiquidText>Raum: {{ lesson.room }}</LiquidText>
            </LiquidCard>
          </LiquidOutline>
        </TabPanel>
      </TabPanels>
    </TabGroup>
  </LiquidPage>
</template>
<script setup>
import {Tab, TabGroup, TabList, TabPanel, TabPanels} from "@headlessui/vue";

const { timetable } = await $fetch("/api/timetable")
const { teachers } = await $fetch("/api/teachers")

function getTeacherByShort(short) {
  for (let i = 0; i < teachers.length; i++) {
    if(teachers[i].short === short) {
      return teachers[i]
    }
  }
  return {}
}

function getCurrentDayOfWeek() {
  const today = new Date().getDay();
  const currentDay = (today + 6) % 7
  return currentDay >= 5 ? 0 : currentDay
}
</script>