<template>
  <Page :title="props.error.statusMessage.split(':')[0]">
    <Wrapper>
      <div class="flex items-center justify-center text-center flex-col my-6">
        <h1 class="text-blue-700 font-bold text-5xl mb-3">{{ props.error.statusCode }}</h1>
        <h1 class="font-bold">{{ props.error.statusMessage.split(':')[0] }}</h1>
        <span>Ein Stacktrace befindet sich unten.</span>
        <span class="m-6 my-12 text-zinc-400 text-sm">{{ props.error.stack }}</span>
      </div>
    </Wrapper>
  </Page>
</template>

<script setup>
const props = defineProps({error: {}})
</script>