<template>
  <UnityPage page-title="Kalender">
    <TabGroup>
      <TabList class="flex gap-1.5 p-1.5 m-3 rounded-lg bg-zinc-100">
        <template v-for="month in exams">
          <Tab v-if="month.exams.length !== 0" class="text-sm outline-none ui-selected:bg-white font-medium ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">{{ month.name }}</Tab>
        </template>
      </TabList>

      <TabPanels>
        <template v-for="month in exams">
          <TabPanel v-if="month.exams.length !== 0">
            <Outline>
              <Card v-for="exam in month.exams">
                <template #header>
                  <h1 class="font-semibold">{{ exam.date }}. {{ month.name }}. {{ month.year }} &middot; <span class="font-normal">{{ exam.name }}</span></h1>
                </template>

                <p>
                  Themen: ...
                </p>

                <template v-if="exam.teacher.length !== 0" #footer>
                  <span class="text-sm text-zinc-400 flex gap-1">
                    <span v-for="teacher in exam.teacher">{{ teacher }}</span>
                  </span>
                </template>
              </Card>
            </Outline>
          </TabPanel>
        </template>
      </TabPanels>
    </TabGroup>
  </UnityPage>
</template>

<script setup>
import {
  Tab,
  TabGroup,
  TabList,
  TabPanel,
  TabPanels
} from "@headlessui/vue";

const date = ref(new Date())
const exams = await getExams()

async function getExams() {
  const { exams } = await $fetch("/api/exams")
  return exams
}

function getDate(d) {
  const date = new Date(d)
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
</script>