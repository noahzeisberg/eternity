<template>
  <Page>
    <Tabs default-value="class">
      <TabsList class="w-full">
        <TabsTrigger value="class" class="w-full">Klassenbasiert</TabsTrigger>
        <TabsTrigger value="full" class="w-full">Volle Ansicht</TabsTrigger>
      </TabsList>

      <TabsContent value="class">
        <Card>
          <CardHeader>
            <CardTitle>Vertretungsplan</CardTitle>
            <CardDescription>Bitte wäle deine Klasse aus, um akkurate Informationen zu erhalten.</CardDescription>
          </CardHeader>

          <CardContent>
            <Input v-model="preferredClass"></Input>
            <CardDescription class="text-center mt-2">Angaben ohne Gewähr.</CardDescription>
          </CardContent>
        </Card>

        <template v-for="substitution in substitutions">
          <Card v-if="substitution.class.includes(preferredClass)">
            <CardHeader>
              {{ removeHTMLTag(substitution.hour) }}
              {{ removeHTMLTag(substitution.class) }}
              {{ removeHTMLTag(substitution.subject) }}
              {{ removeHTMLTag(substitution.room) }}
              {{ removeHTMLTag(substitution.teacher) }}
              {{ removeHTMLTag(substitution.type === "" ? "Vertretung" : substitution.type) }}
            </CardHeader>
          </Card>
        </template>
      </TabsContent>

      <TabsContent value="full" class="flex flex-col gap-3">
        <Card v-for="substitution in substitutions">
          <CardHeader>
            {{ removeHTMLTag(substitution.hour) }}
            {{ removeHTMLTag(substitution.class) }}
            {{ removeHTMLTag(substitution.subject) }}
            {{ removeHTMLTag(substitution.room) }}
            {{ removeHTMLTag(substitution.teacher) }}
            {{ removeHTMLTag(substitution.type === "" ? "Vertretung" : substitution.type) }}
          </CardHeader>
        </Card>
        <h1 class="text-sm text-zinc-500 dark:text-zinc-400 text-center">Angaben ohne Gewähr.</h1>
      </TabsContent>
    </Tabs>
  </Page>
</template>

<script setup lang="ts">
import {useLocalStorage} from "@vueuse/core";
import {Tabs, TabsList, TabsContent, TabsTrigger} from "~/components/ui/tabs";

const preferredClass = useLocalStorage("preferredClass", "8G2")
const { data: substitutions } = useLazyAsyncData("substitutions", () => $fetch("/api/substitutions"))

const removeHTMLTag = (input: string) => {
  return input.replace(/<\/?[^>]+>/g, "");
}
</script>