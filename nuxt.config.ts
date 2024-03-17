// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    devtools: {
        enabled: false
    },
    modules: ['@nuxtjs/tailwindcss', "@nuxt/image", '@nuxtjs/supabase'],
    app: {
        head: {
            title: "Eternity Web - Active Development Preview",
            link: [
                { rel: 'icon', type: 'image/ico', href: '/eternity.png' },
            ],
            meta: [
                { name: "referrer", content: "no-referrer" }
            ]
        },
    },
    ssr: false,
    nitro: {
      storage: {
          database: {
              driver: "fs",
              base: "database"
          }
      }
    },
    routeRules: {
        "/substitutions": {
            cors: true
        }
    }
})