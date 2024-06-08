<template>
  <Page>
    <Tabs :default-value="0">
      <TabsList class="w-full">
        <TabsTrigger class="w-full" v-for="(day, index) in timetable" :value="index">{{ day.day.substring(0, 2) }}</TabsTrigger>
      </TabsList>
      <TabsContent v-for="(day, index) in timetable" :value="index">
        <div class="flex flex-col gap-3">
          <Card v-for="(lesson, index) in day.lessons">
            <CardHeader>
              <CardTitle>{{ index+1 }}. {{ lesson.name }}</CardTitle>
              <CardDescription><TeacherProfile :short="lesson.teacher"></TeacherProfile> in {{ lesson.room }}</CardDescription>
            </CardHeader>
          </Card>
        </div>
      </TabsContent>
    </Tabs>
  </Page>
</template>

<script setup>
import TeacherProfile from "~/components/widgets/TeacherProfile.vue";

const { data: timetable } = useLazyAsyncData("timetable", () => $fetch("/api/timetable"))
</script>