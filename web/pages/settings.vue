<template>
  <Page>
    <Card>
      <CardHeader>
        <CardTitle>Allgemein</CardTitle>
        <CardDescription>Test</CardDescription>
      </CardHeader>
      <CardFooter class="flex-col">
        <div class="flex justify-between w-full">
          <Label for="enabled"></Label>
          <Switch></Switch>
        </div>
      </CardFooter>
    </Card>
  </Page>
</template>
<script setup lang="ts">
import {Switch} from "~/components/ui/switch";
</script>