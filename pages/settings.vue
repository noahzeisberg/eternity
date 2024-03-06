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
        <AboutProfile display-name="Luca Peter" task="Datenverarbeitung & Ideen" picture="luca"/>
        <AboutProfile display-name="Dominik Bauer" task="Datenverarbeitung & Ideen" picture="dominik"/>
        <AboutProfile display-name="Massih Haschemi" task="Datensammlung & Ideen" picture="massih"/>
        <AboutProfile display-name="Elias Wardak" task="Ideen" picture="elias"/>
      </LiquidCard>

      <LiquidCard>
        <template #header>
          <LiquidTitle>App-Informationen</LiquidTitle>
        </template>
        <LiquidText>Version: master@<NuxtLink :to="commitLink" class="text-zinc-400 underline">({{ commitSha }})</NuxtLink></LiquidText>
        <LiquidText>Letztes Update: {{ commitTimeString }}</LiquidText>
        <template #footer>
          <NuxtLink to="https://github.com/noahzeisberg/unity" class="text-zinc-400">GitHub Repository</NuxtLink>
        </template>
      </LiquidCard>

      <NuxtLink to="mailto://noah.zeisberg@heinrichboell.schule" class="text-center w-full py-3 bg-blue-700 text-zinc-50 rounded-lg shadow">Probleme melden</NuxtLink>
    </LiquidOutline>
  </LiquidPage>
</template>

<script setup>
let latestCommit
let commitDate
let commitTimeString
let commitLink
let commitSha

try {
  latestCommit = await $fetch("https://api.github.com/repos/noahzeisberg/unity/commits")[0]
  commitDate = new Date(latestCommit["commit"]["committer"]["date"])
  commitTimeString = commitDate.getDate() + "." + commitDate.getMonth() + "." + commitDate.getFullYear()
  commitLink = latestCommit["html_url"]
  commitSha = latestCommit["sha"].substring(0, 7)
} catch (e) {
  console.log(e)
}
</script>