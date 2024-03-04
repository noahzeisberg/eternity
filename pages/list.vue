<template>
  <Page page-title="Liste">
    <TabGroup>
      <TabList class="flex gap-1.5 p-1.5 m-3 rounded-lg bg-zinc-100">
        <Tab class="text-sm outline-none ui-selected:bg-white font-medium ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Schüler</Tab>
        <Tab class="text-sm outline-none ui-selected:bg-white font-medium ui-selected:text-blue-700 w-full py-2 rounded-md ui-selected:shadow">Lehrer</Tab>
      </TabList>

      <TabPanels>
        <TabPanel>
          <Outline>
            <Card v-for="student in students">
              <template #header>
                <h1 class="font-semibold">{{ student }}</h1>
              </template>

              <template #footer>
                <span class="text-sm text-zinc-400">Anwesend</span>
              </template>
            </Card>
          </Outline>
        </TabPanel>

        <TabPanel>
          <Outline>
            <Card>
              <span class="flex gap-3 items-center">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
                </svg>
                <form class="w-full" @submit.prevent="openTeacherProfile">
                  <input placeholder="Suche nach Kürzel" v-model="input" class="bg-transparent text-center w-full placeholder:text-sm focus:placeholder-transparent outline-none" type="text">
                </form>
                <button @click="openTeacherProfile">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 4.5 21 12m0 0-7.5 7.5M21 12H3" />
                  </svg>
                </button>
              </span>
            </Card>

            <Card v-for="teacher in teachers">
              <template #header>
                <h1 class="font-semibold">{{ teacher["name"] }} <span class="font-normal text-zinc-400">({{ teacher["short"] }})</span></h1>
              </template>

              <span class="text-sm" v-for="subject in teacher.subjects">{{ subject }}</span>
            </Card>
          </Outline>
        </TabPanel>
      </TabPanels>
    </TabGroup>
  </Page>
</template>
<script setup>
import {
  Tab,
  TabGroup,
  TabList,
  TabPanel,
  TabPanels
} from "@headlessui/vue";

const input = ref("")
const { teachers } = await $fetch("/api/teachers")
const { students } = await $fetch("/api/students")

async function openTeacherProfile() {
  await useRouter().push("/teacher/" + input.value.toLowerCase())
}
</script>