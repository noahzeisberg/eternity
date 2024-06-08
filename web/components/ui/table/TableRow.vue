<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'

const props = defineProps<{
  class?: HTMLAttributes['class']
}>()
</script>

<template>
  <tr :class="cn('border-b border-zinc-200 dark:border-zinc-800', props.class)">
    <slot />
  </tr>
</template>
