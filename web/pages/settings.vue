<template>
  <Page>
    <Card>
      <CardHeader>
        <CardTitle>Profil</CardTitle>
        <CardDescription>{{ user?.email }}</CardDescription>
      </CardHeader>
      <CardContent class="flex flex-col gap-2">
        <Button @click="supabase.auth.signOut" class="w-full">Abmelden</Button>
      </CardContent>
    </Card>

    <Card>
      <CardHeader>
        <CardTitle>Mitwirkende</CardTitle>
        <CardDescription>Eine Liste des Kernteams und wichtiger Personen bezüglich der Entwicklung.</CardDescription>
      </CardHeader>
      <CardContent class="flex flex-col gap-1">
        <Contributor name="Noah" role="Core Developer" link="https://noahzeisberg.github.io"></Contributor>
        <Contributor name="Luca" role="Database Operations" link="https://steamcommunity.com/profiles/76561199210639380"></Contributor>
        <Contributor name="Dominik" role="PR Manager & Technical Writer" link="https://github.com/medomi106"></Contributor>
        <Contributor name="Massih" role="Quality Assurance" link=""></Contributor>
        <Contributor name="Raziel" role="Technical Writer" link="https://"></Contributor>
      </CardContent>
    </Card>
  </Page>
</template>

<script setup lang="ts">
import {Switch} from "~/components/ui/switch";
import Contributor from "~/components/widgets/Contributor.vue";

const supabase = useSupabaseClient()
const user = useSupabaseUser()
</script>