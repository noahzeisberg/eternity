<template>
  <UnityPage page-title="Einstellungen">
    <Outline>
      <Card>
        <template #header>
          <h1 class="font-semibold">Präferenzen</h1>
        </template>
        Optionen werden bald hinzugefügt.
      </Card>

      <Card>
        <template #header>
          <h1 class="font-semibold">Erscheinungsbild</h1>
        </template>
        Optionen werden bald hinzugefügt.
      </Card>

      <Card>
        <template #header>
          <h1 class="font-semibold">Verhalten</h1>
        </template>
        Optionen werden bald hinzugefügt.
      </Card>

      <Card>
        <template #header>
          <h1 class="font-semibold">Contributors</h1>
        </template>
        <AboutProfile display-name="Noah Zeisberg" task="Core Developer" picture="noah"/>
        <AboutProfile display-name="Luca Peter" task="Datenverarbeitung & Ideen" picture="luca"/>
        <AboutProfile display-name="Dominik Bauer" task="Datenverarbeitung & Ideen" picture="dominik"/>
        <AboutProfile display-name="Massih Haschemi" task="Ideen & Datenerarbeitung" picture="massih"/>
        <AboutProfile display-name="Elias Wardak" task="Ideen" picture="elias"/>
      </Card>

      <Card>
        <template #header>
          <h1 class="font-semibold">App-Informationen</h1>
        </template>
        <p>Version: master@<NuxtLink :to="commitLink" class="text-zinc-400 underline">({{ commitSha }})</NuxtLink></p>
        <p>Letztes Update: {{ commitTimeString }}</p>
        <template #footer>
          <NuxtLink to="https://github.com/noahzeisberg/unity" class="text-zinc-400">GitHub Repository</NuxtLink>
        </template>
      </Card>

      <NuxtLink to="mailto://noah.zeisberg@heinrichboell.schule" class="text-center w-full py-3 bg-blue-700 text-zinc-50 rounded-lg shadow">Probleme melden</NuxtLink>
    </Outline>
  </UnityPage>
</template>

<script setup>
let latestCommit
let commitDate
let commitTimeString
let commitLink
let commitSha

try {
  latestCommit = await retrieveLatestCommit()
  commitDate = new Date(latestCommit["commit"]["committer"]["date"])
  commitTimeString = commitDate.getDate() + "." + commitDate.getMonth() + "." + commitDate.getFullYear()
  commitLink = latestCommit["html_url"]
  commitSha = latestCommit["sha"].substring(0, 7)

  async function retrieveLatestCommit() {
    const data = await $fetch("https://api.github.com/repos/noahzeisberg/unity/commits")
    return data[0]
  }
}
catch (e) {
  console.log(e)
}
</script>