<template>
  <LiquidPage page-title="Liste">
    <TabGroup>
      <TabList class="flex gap-1.5 p-1.5 m-3 rounded-lg bg-zinc-100">
        <Tab class="text-sm outline-none ui-selected:bg-white font-medium ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Schüler</Tab>
        <Tab class="text-sm outline-none ui-selected:bg-white font-medium ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Lehrer</Tab>
      </TabList>

      <TabPanels>
        <TabPanel>
          <LiquidOutline>
            <LiquidCard v-for="student in students">
              <template #header>
                <LiquidTitle>{{ student }}</LiquidTitle>
              </template>

              <template #footer>
                <LiquidDisclaimer accent>Anwesend</LiquidDisclaimer>
              </template>
            </LiquidCard>
          </LiquidOutline>
        </TabPanel>

        <TabPanel>
          <LiquidOutline>
            <LiquidCard>
              <span class="flex gap-3 items-center">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
                </svg>
                <input placeholder="Suchen..." v-model="query" class="bg-transparent text-center w-full placeholder:text-sm focus:placeholder-transparent outline-none" type="text">
                <div class="w-5 h-5">
                  <button v-if="query !== ''" @click="query = ''">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12"/>
                    </svg>
                  </button>
                </div>
              </span>
            </LiquidCard>

            <LiquidCard v-for="teacher in filteredTeachers">
              <template #header>
                <h1 class="font-semibold">{{ teacher["name"] }} <span class="font-normal text-zinc-400">({{ teacher["short"] }})</span></h1>
              </template>

              <LiquidText v-for="subject in teacher.subjects">{{ subject }}</LiquidText>
            </LiquidCard>
          </LiquidOutline>
        </TabPanel>
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

const query = ref("")
const { teachers } = await $fetch("/api/teachers")
const { students } = await $fetch("/api/students")

const filteredTeachers = computed(() => query.value === '' ? teachers : teachers.filter((teacher) => {
  return teacher.name.toLowerCase().includes(query.value.toLowerCase()) || teacher.short.toLowerCase().includes(query.value.toLowerCase())
}))
</script>