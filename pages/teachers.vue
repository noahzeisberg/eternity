<template>
  <Page title="Lehrkräfte">
    <Wrapper>
      <UInput icon="i-heroicons-magnifying-glass" placeholder="Suche..." v-model="query"></UInput>
    </Wrapper>

    <Loader v-if="pending"></Loader>

    <UTable v-else :columns="columns" :rows="filteredTeachers">
      <template #subjects-data="{ row }">
        <div class="flex flex-col">
          <span v-for="item in row.subjects">{{ item }}</span>
        </div>
      </template>

      <template #contact-data="{ row }">
        <div class="flex gap-1">
          <UButton :to="`mailto://${row.short.toLowerCase()}@hbs-hattersheim.de`" color="gray" icon="i-heroicons-envelope" variant="solid"></UButton>
        </div>
      </template>
    </UTable>
  </Page>
</template>

<script setup lang="ts">
const query = ref("")
const { pending, data: teachers } = await useLazyAsyncData("teachers", () => $fetch('/api/teachers'))
let filteredTeachers

watch(teachers, (list) => {
  if(list === null) return
  filteredTeachers = computed(() => query.value === '' ? list : list.filter((teacher: { name: string, short: string }) => {
    return teacher.name.toLowerCase().includes(query.value.toLowerCase()) || teacher.short.toLowerCase().includes(query.value.toLowerCase())
  }))
})

console.log(teachers)
console.log(teachers.value)

const columns = [
  {
    key: "short",
    label: "Kürzel"
  },
  {
    key: "name",
    label: "Name"
  },
  {
    key: "subjects",
    label: "Fächer"
  },
  {
    key: "contact",
    label: "Mail"
  }
]
</script>