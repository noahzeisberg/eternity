<template>
  <Page>
    <Tabs default-value="today">
      <TabsList class="w-full">
        <TabsTrigger class="w-full" value="today">Heute</TabsTrigger>
        <TabsTrigger class="w-full" value="tomorrow">Morgen</TabsTrigger>
      </TabsList>
      <TabsContent value="today">
        <div class="flex flex-col gap-3">
          <SubstitutionDisplay :substitutions="today"></SubstitutionDisplay>
        </div>
      </TabsContent>

      <TabsContent value="tomorrow">
        <div class="flex flex-col gap-3">
          <SubstitutionDisplay :substitutions="tomorrow"></SubstitutionDisplay>
        </div>
      </TabsContent>
    </Tabs>
    <h1 class="text-sm text-zinc-500 dark:text-zinc-400 text-center">Angaben ohne Gewähr.</h1>
  </Page>
</template>

<script setup lang="ts">
let today = await requestSubstitution(new Date())
let tomorrow = await requestSubstitution(new Date(Date.now() + 24 * 60 * 60 * 1000))

async function requestSubstitution(date: Date) {
  return await $fetch<{
    date: string,
    lastUpdate: string,
    rows: {
      hour: string,
      class: string,
      subject: string,
      room: string,
      teacher: string,
      type: string
    }[],
  }>("/api/substitutions?date=" + Date.parse(date.toString()))
}
</script>