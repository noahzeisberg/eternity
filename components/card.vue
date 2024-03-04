<template>
  <div class="bg-zinc-50 rounded-lg w-full shadow">
    <div class="flex flex-col divide-zinc-200 divide-y w-full">
      <div v-if="$slots.header" class="flex flex-col px-6 py-2">
        <slot name="header"></slot>
      </div>

      <div v-if="$slots.default" class="flex flex-col px-6 py-2">
        <slot></slot>
      </div>

      <div v-if="$slots.footer" class="flex flex-col px-6 py-2">
        <slot name="footer"></slot>
      </div>
    </div>
  </div>
</template>