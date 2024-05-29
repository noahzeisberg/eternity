// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  modules: ["@nuxtjs/tailwindcss", "shadcn-nuxt", "@nuxtjs/supabase"],
  app: {
      head: {
        title: "Eternity — School Organisation System",
        meta: [{ name: "referrer", content: "no-referrer" }]
      },
    },
    ssr: false
})