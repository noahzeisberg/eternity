<template>
  <LiquidPage page-title="Vertretungsplan">
    <LiquidOutline>
      <LiquidCard>
        <template #header>
          <LiquidTitle>Vertretungen für {{ statistics.weekDay }}</LiquidTitle>
        </template>
        Zuletzt aktualisiert: {{ statistics.lastUpdate }}
        <template #footer>
          <LiquidDisclaimer accent>Alle Angaben ohne Gewähr.</LiquidDisclaimer>
        </template>
      </LiquidCard>

      <FormInput class="animation-ready" placeholder="Klasse auswählen..." v-model="preferredClass"></FormInput>

      <LiquidCard v-for="substitution in substitutions">
        <template #header>
          <LiquidTitle>{{ substitution.lessons }}. Stunde <LiquidText accent>({{ substitution.subject }})</LiquidText></LiquidTitle>
        </template>

        <LiquidText>Raum: {{ sanitize(substitution.room) }}</LiquidText>
        <LiquidText>Lehrer: {{ sanitize(substitution.teachers) }}</LiquidText>
        <LiquidText>Typ: {{ substitution.type ? substitution.type : "Vertretung" }}</LiquidText>

        <template #footer>
          <LiquidDisclaimer accent>{{ substitution.classes }}</LiquidDisclaimer>
        </template>
      </LiquidCard>

      <LiquidCard v-if="substitutions.length === 0">
        <LiquidTitle>Keine Vertretungen</LiquidTitle>
      </LiquidCard>
    </LiquidOutline>
  </LiquidPage>
</template>

<script setup>
const preferredClass = ref(useCookie("user.preferences.class").value.trim())
watch(preferredClass, () => {
  useCookie("user.preferences.class").value = preferredClass.value
}, { immediate: true })

const today = new Date()

const { payload } = await getPayload()

const statistics = {
  "date": payload["date"],
  "affectedClasses": payload["affectedElements"][0],
  "lastUpdate": payload["lastUpdate"],
  "weekDay": payload["weekDay"],
}

let subList = []
payload["rows"].forEach((item) => {
  subList.push({
    "lessons": item["data"][0],
    "classes": item["data"][1],
    "subject": item["data"][2],
    "room": item["data"][3],
    "teachers": item["data"][4],
    "type": item["data"][5],
    "additional": item["data"][6],
  })
})

const substitutions = computed(() => preferredClass.value === "" ? subList : subList.filter((substitution) => substitution["classes"].toLowerCase().includes(preferredClass.value.toLowerCase())))

function sanitize(string) {
  return string.replace(/<[^>]*>?/gm, "");
}

function formatDate(inputDate) {
  return inputDate.getFullYear().toString()+pad(inputDate.getMonth()+1, 2).toString()+pad(inputDate.getDate(), 2).toString()
}

function pad(number, length) {
  let str = '' + number;
  while (str.length < length) {
    str = '0' + str;
  }
  return str;
}

async function getPayload() {
  return await $fetch("https://ajax.webuntis.com/WebUntis/monitor/substitution/data?school=hbs-hattersheim", {
    method: "POST",
    body: {
      "formatName": "Vetr_Kla",
      "schoolName": "hbs-hattersheim",
      "date": formatDate(today),
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
      "activityTypeIds": [2,3,4,8],
      "showEvent": true,
      "showCancel": true,
      "showOnlyCancel": false,
      "showSubstTypeColor": false,
      "showExamSupervision": false,
      "showUnheraldedExams": false
    }
  })
}
</script>