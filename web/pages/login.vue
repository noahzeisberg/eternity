<template>
  <div class="w-screen h-screen bg-white dark:bg-zinc-950 flex flex-col items-center justify-center">
    <div class="flex flex-col gap-3 w-full max-w-80">
      <div class="w-full flex justify-center">
        <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="size-12">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M11 8L3 16H11L11 8ZM13 16L21 8L13 8V16Z" fill="currentColor"/>
        </svg>
      </div>
      <div>
        <Label for="email">E-Mail</Label>
        <Input v-model="email" id="email" type="email" />
      </div>
      <div>
        <Label for="password">Passwort</Label>
        <Input v-model="password" id="password" type="password" />
      </div>
      <Button @click="login">Anmelden</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from 'vue-sonner'

const email = ref("")
const password = ref("")

const supabase = useSupabaseClient()

const login = async () => {
  const { error } = await supabase.auth.signInWithPassword({
    email: email.value,
    password: password.value,
  })
  if (error) {
    toast.error(error.message === "Invalid login credentials" ? "Ungültige Anmeldedaten!" : error.message)
  }
}

const user = useSupabaseUser()
watch(user, () => {
  if (user.value?.email) {
    const nameParts = user.value.email.split("@")[0].split(".")
    toast.success("Erfolgreich als " + nameParts[0].replace(/\b\w/, (c) => c.toUpperCase()) + " " + nameParts[1].replace(/\b\w/, (c) => c.toUpperCase()) + " eingeloggt!")
    return navigateTo("/")
  }
})
</script>