// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  modules: ["@nuxtjs/tailwindcss", "shadcn-nuxt", "@nuxtjs/supabase"],
  app: {
      head: {
        title: "Eternity — School Organisation System",
        link: [{ rel: 'icon', type: 'image/svg+xml', href: "data:image/svg+xml,%3Csvg viewBox='0 0 24 24' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath fill-rule='evenodd' clip-rule='evenodd' d='M3.5 16L11.5 8L11.5 16H9.68182V12.3636L6.04545 16H3.5Z' fill='white'/%3E%3Cpath fill-rule='evenodd' clip-rule='evenodd' d='M20.5 8L12.5 16L12.5 8L14.3182 8L14.3182 11.6364L17.9545 8L20.5 8Z' fill='white'/%3E%3C/svg%3E" }],
        meta: [{ name: "referrer", content: "no-referrer" }]
      },
    },
    ssr: false
})