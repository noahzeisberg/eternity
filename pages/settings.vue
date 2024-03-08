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
        <AboutProfile display-name="Noah Zeisberg" task="Core Developer" picture="noah" mail="noahonfyre@gmail.com"/>
        <AboutProfile display-name="Elias Wardak" task="PR Manager & Testing" picture="elias" mail="elias.wardak@gmx.de"/>
        <AboutProfile display-name="Dominik Bauer" task="Marketing & Ideen" picture="dominik" mail="dominik.bauer.106@gmail.com"/>
        <AboutProfile display-name="Massih Haschemi" task="Datenerarbeitung" picture="massih" mail="Massih.h1@gmail.com"/>
        <AboutProfile display-name="Raziel Otten" task="Ideen" picture="raziel" mail="Ottenraziel3@gmail.com"/>
        <AboutProfile display-name="Enrico Sanfratello" task="Ideen" picture="enrico" mail=""/>
        <AboutProfile display-name="Luca Peter" task="Ideen" picture="luca" mail="milchmax2000@gmail.com"/>
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