<template>
  <div class="flex flex-col justify-between text-center items-stretch p-6 h-screen">
    <div class="flex flex-col mt-16">
      <LiquidTitle class="text-2xl font-bold">Eternity Login</LiquidTitle>
      <LiquidDisclaimer accent>Active Development Preview</LiquidDisclaimer>
    </div>

    <div class="flex flex-col gap-3">
      <div class="flex flex-col gap-1">
        <div class="flex justify-between items-center">
          <LiquidText>Benutzer</LiquidText>
          <LiquidDisclaimer accent></LiquidDisclaimer>
        </div>
        <FormInput @click="resetMsg" type="text" v-model="username"></FormInput>
      </div>
      <div class="flex flex-col gap-1">
        <div class="flex justify-between items-center">
          <LiquidText>Passwort</LiquidText>
          <LiquidDisclaimer @click="sendPasswordRecovery" accent><button>Passwort vergessen?</button></LiquidDisclaimer>
        </div>
        <FormInput @click="resetMsg" type="password" v-model="password"></FormInput>
      </div>
      <div v-if="msg !== ''" class="text-sm bg-red-700/15 ring-1 ring-red-700/30 py-2.5 rounded-md">{{ msg }}.</div>
    </div>

    <div class="flex flex-col gap-3">
      <FormButton @click="login">Anmelden</FormButton>
      <FormButton @click="signUp" accent>Registrieren</FormButton>
      <NuxtLink to="mailto://noah.zeisberg@heinrichboell.schule">
        <LiquidDisclaimer accent>Probleme beim Anmelden?</LiquidDisclaimer>
      </NuxtLink>
    </div>
  </div>
</template>

<script setup>
import Button from "~/components/form/button.vue";

const supabase = useSupabaseClient()
const username = ref("")
const password = ref("")
const msg = ref("")
const user = useSupabaseUser()

watch(user, () => {
  if (user.value) {
    return navigateTo("/")
  }
}, { immediate: true })

async function login() {
  const { email, isTeacherLogin } = parseMailAddress(username.value)
  const { error } = await supabase.auth.signInWithPassword({
    email: email,
    password: password.value
  })
  if (error) msg.value = error.message
}

async function signUp() {
  const { email, isTeacherLogin } = parseMailAddress(username.value)
  const { error } = await supabase.auth.signUp({
    email: email,
    password: password.value
  })
  msg.value = "Bitte überprüfe deinen Posteingang"
  if (error) msg.value = error.message
}

async function sendPasswordRecovery() {
  const { email, isTeacherLogin } = parseMailAddress(username.value)
  const { error } = await supabase.auth.resetPasswordForEmail(email)
  msg.value = "Bitte überprüfe deinen Posteingang"
  if (error) msg.value = error.message
}

async function resetMsg() {
  msg.value = ""
}

function parseMailAddress(username) {
  let email
  let isTeacherLogin = false
  if(username.includes("@")) {
    email = username
  } else if(username.includes(".")) {
    email = username + "@heinrichboell.schule"
    isTeacherLogin = false
  } else {
    email = username + "@hbs-hattersheim.de"
    isTeacherLogin = true
  }
  return { email, isTeacherLogin }
}
</script>