<template>
  <Page title="Einstellungen">
    <Wrapper>
      <UCard>
        <template #header>
          <h1 class="font-semibold -my-2">Klasse</h1>
        </template>
        <div class="flex flex-col gap-3">
          <UButtonGroup class="w-full">
            <UInput class="w-full" placeholder="Klasse eingeben." v-model="classObject.className"/>
            <UButton @click="updatePreferences" color="black">Speichern</UButton>
          </UButtonGroup>

          <div class="flex items-center justify-between">
            <UTooltip text="Zeigt die Vertretungen von jeder Klasse an."><h1 class="text-sm">Alle Vertretungen anzeigen</h1></UTooltip>
            <UToggle :ui="{ active: 'bg-black dark:bg-white' }" v-model="classObject.showAll"></UToggle>
          </div>
        </div>
      </UCard>

      <UCard>
        <template #header>
          <h1 class="font-semibold -my-2">Einstellungen</h1>
        </template>
        <div class="flex flex-col gap-3">
          <div class="flex justify-between items-center">
            <h1 class="font-medium">Farbschema</h1>
            <USelectMenu :ui="{ base: 'capitalize'}" v-model="$colorMode.preference" :options="['dark', 'light', 'system']"></USelectMenu>
          </div>
          <UButtonGroup class="w-1/2">
            <UButton @click="updatePreferences" block color="black" variant="solid">Speichern</UButton>
            <UButton to="/" block color="white" variant="solid">Verwerfen</UButton>
          </UButtonGroup>
        </div>
      </UCard>

      <UCard>
        <template #header>
          <h1 class="font-semibold -my-2">Unser Team</h1>
        </template>
        <TeamCard name="Noah Zeisberg" task="Core Developer & System Admin" img="noah" email="noah.zeisberg@heinrichboell.schule"></TeamCard>
        <TeamCard name="Dominik Bauer" task="PR & Technical Writer" img="dominik" email="dominik.bauer@heinrichboell.schule"></TeamCard>
        <TeamCard name="Massih Haschemi" task="PR & QA" img="massih" email="massih.haschemi@heinrichboell.schule"></TeamCard>
        <TeamCard name="Elias Wardak" task="Koordinator und QA" img="elias" email="elias.wardak@heinrichboell.schule"></TeamCard>
        <TeamCard name="Luca Peter" task="Database Admin" img="luca" email="luca.peter@heinrichboell.schule"></TeamCard>
        <TeamCard name="Raziel Otten" task="Technical Writer" img="raziel" email="raziel.otten@heinrichboell.schule"></TeamCard>
        <TeamCard name="Enrico Sanfratello" task="Mitwirkender" img="enrico" email="enrico.sanfratello@heinrichboell.schule"></TeamCard>
      </UCard>

      <UCard>
        <template #header>
          <h1 class="font-semibold -my-2">App-Informationen</h1>
        </template>
        <div class="flex flex-col">
          <span>Letztes Update: {{ commit.time }}</span>
          <span>Version: <ULink class="underline font-semibold" :to="commit.link">master@{{ commit.sha }}</ULink></span>
        </div>
        <template #footer>
          <UButton class="flex gap-3 items-center" color="black" variant="solid" to="https://github.com/noahzeisberg/eternity" block>
            <UIcon name="i-octicon-mark-github-16" dynamic></UIcon>
            <span>GitHub Repository</span>
          </UButton>
        </template>
      </UCard>
    </Wrapper>
  </Page>
</template>

<script setup lang="ts">
import type {Ref} from "vue";
import { useStorage } from '@vueuse/core'

const classObject = useStorage("class", {
  className: "",
  showAll: true
})

const commit: Ref<{time: string, link: string, sha: string}> = ref({time: "", link: "", sha: ""})

const updatePreferences = () => {
  useToast().add({
    title: "Änderungen gespeichert!",
    icon: "i-heroicons-check",
  })
}

const { data: commits } = await useLazyAsyncData<[{html_url: string, sha: string, commit: {committer: {date: string}}}]>("latestCommit", () => $fetch("https://api.github.com/repos/noahzeisberg/eternity/commits"))
watch(commits, (list) => {
  if(list === null) return
  let latestCommit = list[0]
  const commitDate = new Date(latestCommit.commit.committer.date)
  commit.value = {
    time: commitDate.getDate() + "." + commitDate.getMonth() + "." + commitDate.getFullYear(),
    link: latestCommit.html_url,
    sha: latestCommit.sha.substring(0, 7)
  }
})
</script>