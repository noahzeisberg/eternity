<template>
  <LiquidPage page-title="Kalender">
    <TabGroup>
      <TabList class="flex gap-1.5 p-1.5 m-3 rounded-lg bg-zinc-100">
        <template v-for="month in exams">
          <Tab v-if="month.exams.length !== 0" class="text-sm outline-none ui-selected:bg-white font-medium ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">{{ month.name }}</Tab>
        </template>
      </TabList>

      <TabPanels>
        <template v-for="month in exams">
          <TabPanel v-if="month.exams.length !== 0">
            <LiquidOutline>
              <LiquidCard v-for="exam in month.exams">
                <template #header>
                  <LiquidTitle>{{ exam.date }}. {{ month.name }} {{ month.year }} &middot; <LiquidText class="font-normal">{{ exam.name }}</LiquidText></LiquidTitle>
                </template>

                <LiquidText v-for="topic in exam.topics">{{ topic.name }}</LiquidText>

                <template v-if="exam.teacher.length !== 0" #footer>
                  <LiquidDisclaimer class="flex gap-1" accent>
                    <LiquidText v-for="teacher in exam.teacher">{{ teacher }}</LiquidText>
                  </LiquidDisclaimer>
                </template>
              </LiquidCard>
            </LiquidOutline>
          </TabPanel>
        </template>
      </TabPanels>
    </TabGroup>
  </LiquidPage>
</template>

<script setup>
import {
  Tab,
  TabGroup,
  TabList,
  TabPanel,
  TabPanels
} from "@headlessui/vue";

const { exams } = await $fetch("/api/exams")
</script>