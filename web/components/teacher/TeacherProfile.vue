<template>
  <Popover>
    <PopoverTrigger class="underline">{{ short }}</PopoverTrigger>
    <PopoverContent v-if="teacher">
      <h1>{{ teacher?.name }} <span class="text-zinc-500 dark:text-zinc-400">({{ teacher?.short }})</span></h1>
      <p class="text-sm text-zinc-500 dark:text-zinc-400">Fächer: {{ teacher?.subjects.join(", ") }}</p>
    </PopoverContent>
    <PopoverContent v-else>
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

const { data } = useLazyAsyncData("teachers", () => $fetch("/api/teachers"))
watch(data, (list) => {
  if(list === null) return
  for(let item of list) {
    if(item.short == props.short) {
      teacher.value = item
    }
  }
})
</script>