<template>
  <nav class="flex items-center justify-between gap-6 px-6 py-4 sticky top-0 bg-zinc-950">
    <div class="flex items-center justify-center w-8 h-8 text-[10px] text-zinc-50 font-bold bg-blue-700 rounded-full">
      {{ getUserIcon() }}
    </div>

    <h1 class="font-medium">{{ pageTitle }}</h1>
    <NuxtLink class="flex items-center justify-center w-8 h-8" to="/settings">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 12a7.5 7.5 0 0 0 15 0m-15 0a7.5 7.5 0 1 1 15 0m-15 0H3m16.5 0H21m-1.5 0H12m-8.457 3.077 1.41-.513m14.095-5.13 1.41-.513M5.106 17.785l1.15-.964m11.49-9.642 1.149-.964M7.501 19.795l.75-1.3m7.5-12.99.75-1.3m-6.063 16.658.26-1.477m2.605-14.772.26-1.477m0 17.726-.26-1.477M10.698 4.614l-.26-1.477M16.5 19.794l-.75-1.299M7.5 4.205 12 12m6.894 5.785-1.149-.964M6.256 7.178l-1.15-.964m15.352 8.864-1.41-.513M4.954 9.435l-1.41-.514M12.002 12l-3.75 6.495" />
      </svg>
    </NuxtLink>
  </nav>
</template>

<script setup>
defineProps({
  pageTitle: String
})
const user = useSupabaseUser()

function getUserIcon() {
  if (user.value !== null) {
    if (user.value.email.endsWith("@heinrichboell.schule")) {
      let names = user.value.email.split("@")[0].split(".")
      let firstname = names[0]
      let lastname = names[names.length-1]
      return firstname[0].toUpperCase() + lastname[0].toUpperCase()
    } else if (user.value.email.endsWith("@hbs-hattersheim.de")) {
      return user.value.email.replaceAll("@hbs-hattersheim.de", "").toUpperCase()
    }
  } else {
    return "??"
  }
}
</script>