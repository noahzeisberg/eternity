<template>
  <LiquidPage page-title="Seite nicht gefunden.">
    <LiquidOutline>
      <LiquidCard>
        <template #header>
          <h1 class="font-semibold">Diese Seite existiert nicht!</h1>
        </template>

        <span>
          Falls du glaubst, dass dies ein Fehler ist, <NuxtLink class="text-blue-700" to="mailto://noah.zeisberg@heinrichboell.schule">kontaktiere uns</NuxtLink>.
        </span>

        <template #footer>
          <NuxtLink to="/" class="text-sm text-zinc-400">Zurück zur Übersicht</NuxtLink>
        </template>
      </LiquidCard>
    </LiquidOutline>
  </LiquidPage>
</template>