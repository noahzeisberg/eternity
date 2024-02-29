<template>
  <section class="flex flex-col justify-center items-center bg-zinc-100">
    <div class="max-w-xl w-full p-6 pb-0">
      <TabGroup>
        <TabList class="flex justify-stretch gap-3 w-full pb-6">
          <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2">Unterricht</Tab>
          <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2">Schüler</Tab>
          <Tab class="bg-zinc-50 rounded-lg w-full shadow px-6 py-2">Dienste</Tab>
        </TabList>
        <TabPanels>
          <TabPanel>
            <DashboardPageLesson/>
          </TabPanel>

          <TabPanel>
            <DashboardPageStudents/>
          </TabPanel>

          <TabPanel>
            <DashboardPageChores/>
          </TabPanel>
        </TabPanels>
      </TabGroup>
    </div>
  </section>
</template>
<script setup lang="ts">
import {
  Tab,
  TabGroup,
  TabList,
  TabPanel,
  TabPanels
} from "@headlessui/vue";
</script>