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
            <Listbox v-model="currentLanguage" v-slot="{ open }">
              <LiquidCard>
                <div class="flex items-center">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                    <path stroke-linecap="round" stroke-linejoin="round" d="m10.5 21 5.25-11.25L21 21m-9-3h7.5M3 5.621a48.474 48.474 0 0 1 6-.371m0 0c1.12 0 2.233.038 3.334.114M9 5.25V3m3.334 2.364C11.176 10.658 7.69 15.08 3 17.502m9.334-12.138c.896.061 1.785.147 2.666.257m-4.589 8.495a18.023 18.023 0 0 1-3.827-5.802" />
                  </svg>
                  <ListboxButton class="bg-transparent text-center w-full placeholder:text-sm focus:placeholder-transparent outline-none">{{ currentLanguage.name }}</ListboxButton>
                  <div class="w-5 h-5 hover:text-blue-700">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 hover:text-blue-700" v-if="!playSounds" @click="playAudio">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.347a1.125 1.125 0 0 1 0 1.972l-11.54 6.347a1.125 1.125 0 0 1-1.667-.986V5.653Z" />
                    </svg>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 hover:text-blue-700" v-else @click="stopAudio">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 5.25v13.5m-7.5-13.5v13.5" />
                    </svg>
                  </div>
                </div>
                <template v-if="open" #footer>
                  <ListboxOptions>
                    <ListboxOption v-for="language in languages" :value="language">{{ language.name }}</ListboxOption>
                  </ListboxOptions>
                </template>
              </LiquidCard>
            </Listbox>

            <LiquidCard class="duration-300 ease-in-out" :class="currentStudent === student && 'ring-4 ring-blue-700'" v-for="student in students">
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

              <template #footer>
                <LiquidDisclaimer accent>
                  <NuxtLink class="flex gap-3 items-center" :to="`mailto://${teacher.short.toLowerCase()}@hbs-hattersheim.de`">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
                    </svg>
                    <LiquidText>E-Mail</LiquidText>
                  </NuxtLink>
                </LiquidDisclaimer>
              </template>
            </LiquidCard>
          </LiquidOutline>
        </TabPanel>
      </TabPanels>
    </TabGroup>
  </LiquidPage>
</template>
<script setup>
import {
  Listbox,
  ListboxButton,
  ListboxOption,
  ListboxOptions,
  Tab,
  TabGroup,
  TabList,
  TabPanel,
  TabPanels
} from "@headlessui/vue";

const playSounds = ref(false)
const currentStudent = ref("")
const query = ref("")
const { teachers } = await $fetch("/api/teachers")
const { students } = await $fetch("/api/students")

const languages = [
  { name: "Deutsch", id: "de" },
  { name: "Französisch", id: "fr" },
  { name: "Spanisch", id: "es" },
  { name: "Niederländisch", id: "nl" },
  { name: "Polnisch", id: "pl" },
  { name: "Russisch", id: "ru" },
  { name: "Japanisch", id: "ja" },
]

const currentLanguage = ref(languages[0])

const filteredTeachers = computed(() => query.value === '' ? teachers : teachers.filter((teacher) => {
  return teacher.name.toLowerCase().includes(query.value.toLowerCase()) || teacher.short.toLowerCase().includes(query.value.toLowerCase())
}))

async function playAudio() {
  playSounds.value = true
  for (const student of students) {
    if (playSounds.value) {
      currentStudent.value = student
      let url = "https://translate.google.com/translate_tts?ie=UTF-8&tl=" + currentLanguage.value.id + "&client=tw-ob&q=" + student
      let audio = new Audio(url)
      await audio.play()
      await new Promise(resolve => { setTimeout(resolve,  2000) })
    }
  }
  await stopAudio()
}

async function stopAudio() {
  playSounds.value = false
  currentStudent.value = ""
}
</script>