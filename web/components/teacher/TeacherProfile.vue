<template>
  <Popover>
    <PopoverTrigger @click="handleOpen" class="underline">{{ short }}</PopoverTrigger>
    <PopoverContent v-if="teacher">
      <h1>
        {{ teacher?.name }}
        <span class="text-zinc-500 dark:text-zinc-400">({{ teacher?.short }})</span>
      </h1>
      <p class="text-sm text-zinc-500 dark:text-zinc-400">Fächer: {{ teacher?.subjects.join(", ") }}</p>
      <NuxtLink class="w-full" :to="'mailto://' + teacher?.short.toLowerCase() + '@hbs-hattersheim.de'">
        <Button variant="outline" class="w-full mt-2">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 mr-2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
          </svg>
          E-Mail senden
        </Button>
      </NuxtLink>
    </PopoverContent>
    <PopoverContent v-else class="w-full">
      <p class="text-sm text-zinc-500 dark:text-zinc-400">Lehrkraft nicht gefunden!</p>
    </PopoverContent>
  </Popover>
</template>

<script setup lang="ts">
import {Popover, PopoverContent, PopoverTrigger} from "~/components/ui/popover";

const teacher = ref<{
  short: string,
  name: string,
  subjects: string[]
}>()

const props = defineProps({
  short: {
    type: String,
    default: "PLC"
  }
})

const handleOpen = async () => {
  const teachers = await $fetch("/api/teachers")
  for(let item of teachers) {
    if(item.short.toLowerCase() == props.short.toLowerCase()) {
      teacher.value = item
    }
  }
}
</script>