<template>
  <Page>
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
  </Page>
</template>

<script setup lang="ts">
const { data: students } = useLazyAsyncData("students", () => $fetch("/api/students"))
</script>