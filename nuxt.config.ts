// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    devtools: {
        enabled: false
    },
    modules: ['@nuxtjs/tailwindcss', "@nuxt/image"],
    app: {
        head: {
            title: "Eternity for Schools - Active Development Preview",
            link: [
                { rel: 'icon', type: 'image/ico', href: '/eternity.png' },
            ],
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
    }
})