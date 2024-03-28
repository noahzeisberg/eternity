<template>
  <Page title="Vertretungsplan">
    <Loader v-if="pending"></Loader>

    <UTable class="w-full" v-else :columns="columns" :rows="substitutions">
      <template #type-data="{ row }"><span v-html="row.type"></span></template>
      <template #class-data="{ row }">
        <div class="flex flex-col gap-1">
          <span v-for="c in row.class.split(',')">{{ c.trim() }}</span>
        </div>
      </template>
      <template #empty-state>
        <div class="mx-auto text-center my-12">Keine Vertretungen!</div>
      </template>
    </UTable>
  </Page>
</template>

<script setup lang="ts">
import { useStorage } from '@vueuse/core'
const classPreferences = useStorage("class", {className: "", showAll: true})
const apiRoute = classPreferences.value.showAll ? "/api/substitutions" : "/api/substitutions?class=" + classPreferences.value.className

const { pending, data: substitutions } = await useLazyAsyncData("substitutions", () => $fetch(apiRoute))

const columns = [
  {
    key: "hour",
    label: "Stunde"
  },
  {
    key: "class",
    label: "Klasse"
  },
  {
    key: "subject",
    label: "Fach"
  },
  {
    key: "room",
    label: "Raum"
  },
  {
    key: "teacher",
    label: "Lehrer"
  },
  {
    key: "type",
    label: "Typ"
  }
]
</script>