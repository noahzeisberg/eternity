<template>
  <Card>
    <CardHeader>
      <CardTitle>Vertretungsplan für den {{ substitutions ? parseDate(substitutions.date) : "" }}</CardTitle>
      <CardDescription>Zuletzt aktualisiert: {{ substitutions?.lastUpdate }}</CardDescription>
    </CardHeader>
  </Card>

  <Card v-for="substitution in substitutions.rows">
    <CardHeader class="flex-row items-center gap-6 px-8">
      <div class="text-xl font-semibold">{{ removeHTMLTag(substitution.hour) }}</div>

      <div>
        <h1 class="font-semibold">
          {{ removeHTMLTag(substitution.type === "" ? "Vertretung" : substitution.type) }}
          <span class="text-zinc-500 dark:text-zinc-400 font-normal">({{ removeHTMLTag(substitution.subject) }})</span>
        </h1>
        <span class="text-zinc-500 dark:text-zinc-400">
              <TeacherProfile :short="removeHTMLTag(substitution.teacher)"></TeacherProfile>
              <span> in {{ removeHTMLTag(substitution.room) }}</span>
          </span>
      </div>
    </CardHeader>
  </Card>
</template>

<script setup lang="ts">
import TeacherProfile from "~/components/widgets/TeacherProfile.vue";

defineProps({
  substitutions: {
    required: true,
    type: Object as () => {
      date: string,
      lastUpdate: string,
      rows: {
        hour: string,
        class: string,
        subject: string,
        room: string,
        teacher: string,
        type: string
      }[]
    },
  }
})

function parseDate(input: string) {
  const year = parseInt(input.toString().substring(0, 4))
  const month = parseInt(input.toString().substring(4, 6))
  const day = parseInt(input.toString().substring(6, 8))
  const d = new Date(year, month, day)
  return d.getDate().toString() + "." + d.getMonth().toString() + "." + d.getFullYear().toString()
}

function removeHTMLTag(input: string) {
  return input.replace(/<\/?[^>]+>/g, "")
}
</script>