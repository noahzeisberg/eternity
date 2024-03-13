<template>Abmelden...</template>

<script setup>
await useSupabaseClient().auth.signOut()
await useRouter().push("/login")
</script>