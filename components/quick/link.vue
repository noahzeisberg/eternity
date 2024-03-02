<template>
  <NuxtLink class="flex flex-col items-center justify-center gap-0.5" :to="to">
    <slot v-if="isActive" name="active"></slot>
    <slot v-else name="inactive"></slot>

    <span class="text-[12px] font-medium -mt-0.5">{{ title }}</span>
  </NuxtLink>
</template>

<script setup>
const props = defineProps({
  title: String,
  to: String
})

const isActive = ref(useRoute().path.trim() === props.to.trim())
</script>