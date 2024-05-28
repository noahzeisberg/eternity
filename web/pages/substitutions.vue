<template>
  <Page>
    <Tabs default-value="today">
      <TabsList class="w-full">
        <TabsTrigger class="w-full" value="today">Heute</TabsTrigger>
        <TabsTrigger class="w-full" value="tomorrow">Morgen</TabsTrigger>
      </TabsList>
      <TabsContent value="today">
        <div class="flex flex-col gap-3">
          <SubstitutionDisplay :substitutions="today"></SubstitutionDisplay>
        </div>
      </TabsContent>

      <TabsContent value="tomorrow">
        <div class="flex flex-col gap-3">
          <SubstitutionDisplay :substitutions="tomorrow"></SubstitutionDisplay>
        </div>
      </TabsContent>
    </Tabs>
    <h1 class="text-sm text-zinc-500 dark:text-zinc-400 text-center">Angaben ohne Gewähr.</h1>
  </Page>
</template>

<script setup lang="ts">
import sanitize from "sanitize-html";

let today = await requestSubstitution(new Date())
let tomorrow = await requestSubstitution(new Date(Date.now() + 24 * 60 * 60 * 1000))

async function requestSubstitution(date: Date) {
  const payload = (await $fetch<{
    payload: { rows: { data: any }[], date: string, lastUpdate: string, weekDay: string }
  }>("https://ajax.webuntis.com/WebUntis/monitor/substitution/data?school=hbs-hattersheim", {
    method: "POST",
    body: {
      "formatName": "Vetr_Kla",
      "schoolName": "hbs-hattersheim",
      "date": formatDate(date),
      "dateOffset": 0,
      "strikethrough": false,
      "mergeBlocks": false,
      "showOnlyFutureSub": false,
      "showBreakSupervisions": false,
      "showTeacher": true,
      "showClass": true,
      "showHour": true,
      "showInfo": true,
      "showRoom": true,
      "showSubject": true,
      "groupBy": 1,
      "hideAbsent": false,
      "departmentIds": [],
      "departmentElementType": -1,
      "hideCancelWithSubstitution": true,
      "hideCancelCausedByEvent": true,
      "showTime": false,
      "showSubstText": false,
      "showAbsentElements": [],
      "showAffectedElements": [1],
      "showUnitTime": true,
      "showMessages": true,
      "showStudentgroup": false,
      "enableSubstitutionFrom": false,
      "showSubstitutionFrom": 1800,
      "showTeacherOnEvent": false,
      "showAbsentTeacher": false,
      "strikethroughAbsentTeacher": false,
      "activityTypeIds": [2, 3, 4, 8],
      "showEvent": true,
      "showCancel": true,
      "showOnlyCancel": false,
      "showSubstTypeColor": false,
      "showExamSupervision": false,
      "showUnheraldedExams": false
    }
  })).payload

  let rows: any[] = []
  payload.rows.forEach(item => {
    rows.push({
      hour: sanitize(item.data[0]),
      class: sanitize(item.data[1]),
      subject: sanitize(item.data[2]),
      room: sanitize(item.data[3]),
      teacher: sanitize(item.data[4]),
      type: sanitize(item.data[5])
    })
  })
  return {
    "date": payload.date,
    "lastUpdate": payload.lastUpdate,
    "rows": rows.filter((item) => item.class.includes("8G2")),
  }
}

function formatDate(inputDate: Date) {
  return inputDate.getFullYear().toString() + pad(inputDate.getMonth() + 1, 2).toString() + pad(inputDate.getDate(), 2).toString();
}

function pad(number: number, length: number) {
  let str = '' + number;
  while (str.length < length) {
    str = '0' + str;
  }
  return str;
}
</script>