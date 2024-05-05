<template>
  <Page>
    <Card>
      <CardHeader>
        <CardTitle>Vertretungsplan für den {{ substitutions ? parseDate(substitutions.date) : "Tag" }}</CardTitle>
        <CardDescription>Bitte wäle deine Klasse aus, um akkurate Informationen zu erhalten.</CardDescription>
      </CardHeader>

      <CardContent>
        <Input v-model="preferredClass"></Input>
        <h1 class="text-sm text-zinc-500 dark:text-zinc-400 text-center mt-2">Zuletzt aktualisiert: {{ substitutions?.lastUpdate }}</h1>
      </CardContent>
    </Card>

    <template v-for="substitution in substitutions?.rows">
      <Card v-if="substitution.class.includes(preferredClass)">
        <CardHeader class="flex-row items-center gap-6 px-8">
          <div class="text-xl font-semibold">{{ removeHTMLTag(substitution.hour) }}</div>

          <div>
            <h1 class="font-semibold">
              {{ removeHTMLTag(substitution.type === "" ? "Vertretung" : substitution.type) }}
              <span class="text-zinc-500 dark:text-zinc-400 font-normal">({{ removeHTMLTag(substitution.subject) }})</span>
            </h1>
            <span>
              <TeacherProfile :short="removeHTMLTag(substitution.teacher)"></TeacherProfile>
              <span> in {{ removeHTMLTag(substitution.room) }}</span>
            </span>
          </div>
        </CardHeader>
      </Card>
    </template>
    <h1 class="text-sm text-zinc-500 dark:text-zinc-400 text-center">Angaben ohne Gewähr.</h1>
  </Page>
</template>

<script setup lang="ts">
import {useLocalStorage} from "@vueuse/core";

const preferredClass = useLocalStorage("preferredClass", "8G2")
const { data: substitutions } = useLazyAsyncData("substitutions", () => $fetch("/api/substitutions"))

const removeHTMLTag = (input: string) => {
  return input.replace(/<\/?[^>]+>/g, "");
}

const parseDate = (input: string) => {
  const year = parseInt(input.toString().substring(0, 4))
  const month = parseInt(input.toString().substring(4, 6)) - 1
  const day = parseInt(input.toString().substring(6, 8))
  const d = new Date(year, month, day)
  return d.getDate().toString() + "." + d.getMonth().toString() + "." + d.getFullYear().toString()
}
</script>