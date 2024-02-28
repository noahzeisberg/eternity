<template>

</template>

<script setup>
const { data, pending, error, refresh } = await useFetch("https://arche.webuntis.com/WebUntis/api/rest/extern/v1/students/{{id}}")
console.log()
</script>