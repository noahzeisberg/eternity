<template>
  <Page>
    <Tabs default-value="students">
      <TabsList class="w-full">
        <TabsTrigger class="w-full" value="students">Schüler</TabsTrigger>
        <TabsTrigger class="w-full" value="teachers">Lehrer</TabsTrigger>
      </TabsList>
      <TabsContent value="students">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="student in students">
              <TableCell>
                {{ student }}
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </TabsContent>

      <TabsContent class="flex flex-col gap-3" value="teachers">
        <Input v-model="teacherSearch" placeholder="Suchen..."></Input>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Kürzel</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>Fächer</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="teacher in filteredTeacher">
              <TableCell>
                <TeacherProfile :short="teacher.short"></TeacherProfile>
              </TableCell>

              <TableCell>{{ teacher.name }}</TableCell>

              <TableCell class="flex flex-col gap-1">
                <div v-for="subject in teacher.subjects">{{ subject }}</div>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </TabsContent>
    </Tabs>
  </Page>
</template>

<script setup lang="ts">
const { data: students } = useLazyAsyncData("students", () => $fetch("/api/students"))
const { data: teachers } = useLazyAsyncData("teachers", () => $fetch("/api/teachers"))

const teacherSearch = ref("")

const filteredTeacher = computed(() => {
  return teachers.value?.filter(teacher => teacher.name.toLowerCase().includes(teacherSearch.value.toLowerCase()) || teacher.short.toLowerCase().includes(teacherSearch.value.toLowerCase()))
})
</script>