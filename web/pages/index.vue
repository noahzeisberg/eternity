<template>
  <Page>
    <Card>
      <CardHeader class="text-center">
        <CardTitle>{{ getCurrentDayOfWeek() }}</CardTitle>
        <CardDescription>{{ getDate() }}</CardDescription>
      </CardHeader>
    </Card>

    <Card>
      <CardHeader>
        <CardTitle>Dienste</CardTitle>
        <CardDescription>
          <template v-for="assigned in chores">{{ assigned }}<br></template>
        </CardDescription>
      </CardHeader>
    </Card>

    <Card>
      <CardHeader>
        <CardTitle>Dienste</CardTitle>
        <CardDescription></CardDescription>
      </CardHeader>
    </Card>
  </Page>
</template>

<script setup lang="ts">
const { data: chores } = useLazyAsyncData("chores", () => $fetch("/api/chores"))

const getDate = () => {
  const date = new Date()
  return date.getDate() + ". " + [
    "Januar",
    "Februar",
    "März",
    "April",
    "Mai",
    "Juni",
    "Juli",
    "August",
    "September",
    "Oktober",
    "November",
    "Dezember"
  ][date.getMonth()] + " " + date.getFullYear()
};

const getCurrentDayOfWeek = () => {
  const today = new Date().getDay();
  const currentDay = (today + 6) % 7
  return [
    "Montag",
    "Dienstag",
    "Mittwoch",
    "Donnerstag",
    "Freitag",
    "Samstag",
    "Sonntag",
  ][currentDay]
}
</script>