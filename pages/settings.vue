<template>
  <UnityPage page-title="Einstellungen">
    <Outline>
      <Card>
        <template #header>
          <h1 class="font-semibold">Contributors</h1>
        </template>
        <AboutProfile display-name="Noah Zeisberg" task="Core Developer" username="noahzeisberg" github/>
        <AboutProfile display-name="Luca Peter" task="Datenverarbeitung & Ideen" username="lucaonfyre" github/>
        <AboutProfile display-name="Dominik Bauer" task="Datenverarbeitung" username="medomi106" github/>
        <AboutProfile display-name="Massih Haschemi" task="Ideen & Maintenance" username="massih"/>
        <AboutProfile display-name="Elias Wardak" task="Ideen" username="elias795" github/>
      </Card>

      <Card>
        <template #header>
          <h1 class="font-semibold">App-Informationen</h1>
        </template>
        <p>Version: master@<NuxtLink :to="commitLink" class="text-zinc-400 underline">({{ commitSha }})</NuxtLink></p>
        <p>Letztes Update: {{ commitTimeString }}</p>
      </Card>

      <NuxtLink to="mailto://noah.zeisberg@heinrichboell.schule" class="text-center w-full py-3 bg-blue-700 text-zinc-50 rounded-lg">Probleme melden</NuxtLink>
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
  commitTimeString = commitDate.getUTCDay() + "." + commitDate.getUTCMonth() + "." + commitDate.getUTCFullYear()
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