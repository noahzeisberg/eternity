<template>
  <div class="w-screen h-screen bg-white dark:bg-zinc-950 flex flex-col items-center justify-center">
    <Loader></Loader>
    <h1>Abmelden...</h1>
  </div>
</template>

<script setup lang="ts">
import {toast} from "vue-sonner";

const logout = async () => {
  toast.info("Abmeldung erfolgt...")
  const { error } = await useSupabaseClient().auth.signOut()
  if (error) {
    toast.error(error)
  } else {
    toast.success("Erfolgreich abgemeldet!")
    navigateTo("/")
  }
}

logout()
</script>