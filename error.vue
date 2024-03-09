<template>
  <LiquidPage :page-title="props.error.statusMessage.split(':')[0]">
    <LiquidOutline>
      <div class="flex items-center justify-center text-center flex-col my-6">
        <LiquidTitle class="text-blue-700 font-bold text-5xl mb-3">{{ props.error.statusCode }}</LiquidTitle>
        <LiquidTitle>{{ props.error.statusMessage.split(':')[0] }}</LiquidTitle>
        <LiquidDisclaimer>Ein Stacktrace befindet sich unten.</LiquidDisclaimer>
        <LiquidDisclaimer class="mx-6 my-12" accent>{{ props.error.stack }}</LiquidDisclaimer>
      </div>
    </LiquidOutline>
  </LiquidPage>
</template>

<script setup>
const props = defineProps({error: {}})
</script>