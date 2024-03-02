<template>
  <UnityPage :page-title="teacherObject.name">
    <Outline>
      <Card>
        <template #header>
          <h1 class="font-bold">{{ teacherObject.name }} <span class="text-zinc-500">({{ teacherObject.short }})</span></h1>
        </template>
        <div v-for="subject in teacherObject.subjects">
          {{ subject }}
        </div>
        <template #footer>
          <NuxtLink class="text-zinc-500 flex items-center gap-2" :to="`mailto://${teacherObject.short.toLowerCase()}@hbs-hattersheim.de`">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
            </svg>
            <span>E-Mail</span>
          </NuxtLink>
        </template>
      </Card>
    </Outline>
  </UnityPage>
</template>

<script setup>
const { teacher } = useRoute().params
const { teachers } = await getTeachers()

const teacherObject = getTeacherByShort(teacher)

if(teacherObject === null) {
  await useRouter().go(-1)
}

function getTeacherByShort(short) {
  for (let i = 0; i < teachers.length; i++) {
    if(teachers[i].short.toLowerCase() === short.toLowerCase()) {
      return teachers[i]
    }
  }
  return null
}

async function getTeachers() {
  const { teachers } = await $fetch("/api/teachers")
  return { teachers }
}
</script>