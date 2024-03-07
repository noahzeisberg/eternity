<template>
  <LiquidPage page-title="Einstellungen">
    <LiquidOutline>
      <LiquidCard>
        <template #header>
          <LiquidTitle>Präferenzen</LiquidTitle>
        </template>
        Optionen werden bald hinzugefügt.
      </LiquidCard>

      <LiquidCard>
        <template #header>
          <LiquidTitle>Erscheinungsbild</LiquidTitle>
        </template>
        Optionen werden bald hinzugefügt.
      </LiquidCard>

      <LiquidCard>
        <template #header>
          <LiquidTitle>Verhalten</LiquidTitle>
        </template>
        Optionen werden bald hinzugefügt.
      </LiquidCard>

      <LiquidCard>
        <template #header>
          <LiquidTitle>Contributors</LiquidTitle>
        </template>
        <AboutProfile display-name="Noah Zeisberg" task="Core Developer" picture="noah"/>
        <AboutProfile display-name="Elias Wardak" task="PR Manager & Testing" picture="elias"/>
        <AboutProfile display-name="Dominik Bauer" task="Marketing & Ideen" picture="dominik"/>
        <AboutProfile display-name="Massih Haschemi" task="Datenerarbeitung" picture="massih"/>
        <AboutProfile display-name="Enrico Sanfratello" task="Ideen" picture="enrico"/>
        <AboutProfile display-name="Luca Peter" task="Ideen" picture="luca"/>
      </LiquidCard>

      <LiquidCard>
        <template #header>
          <LiquidTitle>App-Informationen</LiquidTitle>
        </template>
        <LiquidText>Version: master@<NuxtLink :to="commitLink" class="text-zinc-400 underline">({{ commitSha }})</NuxtLink></LiquidText>
        <LiquidText>Letztes Update: {{ commitTimeString }}</LiquidText>
        <template #footer>
          <NuxtLink to="https://github.com/noahzeisberg/eternity" class="text-zinc-400">GitHub Repository</NuxtLink>
        </template>
      </LiquidCard>

      <NuxtLink to="mailto://noah.zeisberg@heinrichboell.schule" class="text-center w-full py-3 bg-blue-700 text-zinc-50 rounded-lg shadow">Probleme melden</NuxtLink>
    </LiquidOutline>
  </LiquidPage>
</template>

<script setup>
const commits = await $fetch("https://api.github.com/repos/noahzeisberg/eternity/commits")
let latestCommit = commits[0]
let commitDate = new Date(latestCommit["commit"]["committer"]["date"])
let commitTimeString = commitDate.getDate() + "." + commitDate.getMonth() + "." + commitDate.getFullYear()
let commitLink = latestCommit["html_url"]
let commitSha = latestCommit["sha"].substring(0, 7)
</script>