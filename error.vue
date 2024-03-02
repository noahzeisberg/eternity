<template>
  <UnityPage page-title="Seite nicht gefunden.">
    <Outline>
      <Card>
        <template #header>
          <h1 class="font-semibold">Diese Seite existiert nicht!</h1>
        </template>

        <span>
          Falls du glaubst, dass dies ein Fehler ist, <NuxtLink class="text-blue-700" to="mailto://noah.zeisberg@heinrichboell.schule">kontaktiere uns</NuxtLink>.
        </span>

        <template #footer>
          <NuxtLink class="text-sm text-zinc-400">Zurück zur Übersicht</NuxtLink>
        </template>
      </Card>
    </Outline>
  </UnityPage>
</template>

<script setup lang="ts">
import type { NuxtError } from '#app'

const props = defineProps({
  error: Object as () => NuxtError
})
</script>