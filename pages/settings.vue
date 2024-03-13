<template>
  <LiquidPage page-title="Einstellungen">
    <LiquidOutline>
      <LiquidCard>
        <template #header>
          <div class="flex justify-between items-center">
            <LiquidTitle>Dein Konto</LiquidTitle>
            <NuxtLink to="/logout" class="text-blue-700 text-sm font-medium">Abmelden</NuxtLink>
          </div>
        </template>
        <span class="w-full text-center">{{ user.email }}</span>
      </LiquidCard>

      <LiquidCard>
        <template #header>
          <LiquidTitle>Unser Team</LiquidTitle>
        </template>
        <AboutProfile display-name="Noah Zeisberg" task="Core Developer & System Admin" picture="noah" mail="noahonfyre@gmail.com"/>
        <AboutProfile display-name="Dominik Bauer" task="PR Manager & Technical Writer" picture="dominik" mail="dominik.bauer.106@gmail.com"/>
        <AboutProfile display-name="Elias Wardak" task="Coordinator & Quality Assurance" picture="elias" mail="elias.wardak@gmx.de"/>
        <AboutProfile display-name="Luca Peter" task="Database Admin" picture="luca" mail="milchmax2000@gmail.com"/>
        <AboutProfile display-name="Raziel Otten" task="Technical Writer" picture="raziel" mail="Ottenraziel3@gmail.com"/>
        <AboutProfile display-name="Massih Haschemi" task="Quality Assurance" picture="massih" mail="Massih.h1@gmail.com"/>
        <AboutProfile display-name="Enrico Sanfratello" task="Mitwirkender" picture="enrico" mail=""/>
      </LiquidCard>

      <LiquidCard>
        <template #header>
          <LiquidTitle>Open-Source <LiquidText accent>(via NPM)</LiquidText></LiquidTitle>
        </template>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/nuxt">Nuxt</NuxtLink>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/@nuxtjs/tailwindcss">TailwindCSS for Nuxt</NuxtLink>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/vue">Vue.js</NuxtLink>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/vue-router">Vue Router</NuxtLink>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/@headlessui/vue">Headless UI</NuxtLink>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/@nuxt/image">Nuxt Image</NuxtLink>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/jssoup">JSSoup</NuxtLink>
        <NuxtLink class="hover:text-blue-700" to="https://www.npmjs.com/package/@nuxtjs/supabase">Supabase for Nuxt</NuxtLink>
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
    </LiquidOutline>
  </LiquidPage>
</template>

<script setup>
const user = useSupabaseUser()

const commits = await $fetch("https://api.github.com/repos/noahzeisberg/eternity/commits")
let latestCommit = commits[0]
let commitDate = new Date(latestCommit["commit"]["committer"]["date"])
let commitTimeString = commitDate.getDate() + "." + commitDate.getMonth() + "." + commitDate.getFullYear()
let commitLink = latestCommit["html_url"]
let commitSha = latestCommit["sha"].substring(0, 7)
</script>