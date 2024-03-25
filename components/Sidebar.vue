<template>
  <USlideover v-model="isOpen">
    <UCard class="h-full rounded-none">
      <template #header>
        <div class="flex justify-between items-center">
          <h1 class="font-semibold">Eternity</h1>
          <UButton @click="isOpen = false" color="gray" variant="ghost" icon="i-heroicons-x-mark-20-solid"></UButton>
        </div>
      </template>

      <UVerticalNavigation :ui="{
        size: 'text-md',
        padding: 'px-4 py-2'
      }" :links="links" />
    </UCard>
  </USlideover>
</template>

<script setup lang="ts">
const isOpen = defineModel()
const modal = useModal()

const links = [
  [
    {
      label: "Übersicht",
      icon: "i-heroicons-home",
      to: "/"
    },
    {
      label: "Stundenplan",
      icon: "i-heroicons-clock",
      to: "/timetable"
    },
    {
      label: "Vertretungsplan",
      icon: "i-heroicons-users",
      to: "/substitutions"
    },
    {
      label: "Arbeitsplan",
      icon: "i-heroicons-calendar-days",
      to: "/exams"
    },
    {
      label: "Schüler",
      icon: "i-heroicons-user",
      to: "/students"
    },
    {
      label: "Lehrer",
      icon: "i-heroicons-briefcase",
      to: "/teachers"
    },
    {
      label: "Neuigkeiten",
      icon: "i-heroicons-newspaper",
      to: "/feed"
    },
  ], [
    {
      label: "Einstellungen",
      icon: "i-heroicons-cog",
      to: "/settings"
    },
    {
      label: "Weg",
      icon: "i-heroicons-map-pin",
      to: "https://maps.app.goo.gl/ekwavW7uog2tSbfKA"
    }
  ]
]
</script>