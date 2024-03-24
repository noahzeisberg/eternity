<template>
  <Page title="Vertretungsplan">
    <Loader v-if="pending"></Loader>

    <UTable v-else :columns="columns" :rows="substitutions">
      <template #type-data="{ row }"><span v-html="row.type"></span></template>
    </UTable>
  </Page>
</template>

<script setup lang="ts">
const classCookie = useCookie("user.preferences.class")
const { pending, data: substitutions } = await useLazyAsyncData("substitutions", () => $fetch('/api/substitutions?class=' + classCookie.value))

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