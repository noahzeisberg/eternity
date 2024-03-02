<template>
  <nav class="flex items-center justify-between gap-6 px-6 py-4 sticky top-0 bg-zinc-50">
    <div class="flex items-center">
      <button @click="useRouter().go(-1)" v-if="!isHome">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5 8.25 12l7.5-7.5" />
        </svg>
      </button>

      <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="w-8 h-8">
        <path d="M4 16V8H6.19062L9.67625 12.2286V8H12V16H9.94251L6.32375 11.6114V16H4Z" fill="black"/>
        <path d="M12 8L20 8L20 10.1906L15.7714 13.6762L20 13.6762L20 16L12 16L12 13.9425L16.3886 10.3238L12 10.3238L12 8Z" fill="black"/>
      </svg>

    </div>
    <h1 class="font-medium">{{ pageTitle }}</h1>
    <NuxtLink class="flex items-center justify-center w-8 h-8" to="/settings">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 12a7.5 7.5 0 0 0 15 0m-15 0a7.5 7.5 0 1 1 15 0m-15 0H3m16.5 0H21m-1.5 0H12m-8.457 3.077 1.41-.513m14.095-5.13 1.41-.513M5.106 17.785l1.15-.964m11.49-9.642 1.149-.964M7.501 19.795l.75-1.3m7.5-12.99.75-1.3m-6.063 16.658.26-1.477m2.605-14.772.26-1.477m0 17.726-.26-1.477M10.698 4.614l-.26-1.477M16.5 19.794l-.75-1.299M7.5 4.205 12 12m6.894 5.785-1.149-.964M6.256 7.178l-1.15-.964m15.352 8.864-1.41-.513M4.954 9.435l-1.41-.514M12.002 12l-3.75 6.495" />
      </svg>
    </NuxtLink>
  </nav>
</template>

<script setup lang="ts">
defineProps(["pageTitle"])

const isHome = ref(false)
const route = ref(useRoute().path.split("/")[useRoute().path.split("/").length-1])
isHome.value = route.value === ""
</script>