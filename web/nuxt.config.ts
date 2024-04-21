// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  modules: ["@nuxtjs/tailwindcss", "shadcn-nuxt"],
  app: {
      head: {
        title: "Eternity Web - Active Development Preview",
        meta: [{ name: "referrer", content: "no-referrer" }]
      },
    },
})