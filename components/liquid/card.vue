<template>
  <div class="bg-zinc-900 rounded-md w-full shadow animation-ready ring-1 ring-zinc-800">
    <div class="flex flex-col divide-zinc-800 divide-y w-full">
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