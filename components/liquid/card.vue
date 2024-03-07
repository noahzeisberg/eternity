<template>
  <div class="bg-white rounded-lg w-full shadow animation-ready">
    <div class="flex flex-col divide-zinc-100 divide-y w-full">
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